// Package Poloniex is an implementation of the Poloniex API in Golang.
package poloniex

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	API_BASE = "https://poloniex.com" // Poloniex API endpoint
)

// New returns an instantiated poloniex struct
func New(apiKey, apiSecret string) *Poloniex {
	client := NewClient(apiKey, apiSecret)
	return &Poloniex{client}
}

// New returns an instantiated poloniex struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Poloniex {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Poloniex{client}
}

// poloniex represent a poloniex client
type Poloniex struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Poloniex) SetDebug(enable bool) {
	c.client.debug = enable
}

// GetTickers is used to get the ticker for all markets
func (b *Poloniex) GetTickers() (tickers map[string]Ticker, err error) {
	r, err := b.client.do("GET", "public?command=returnTicker", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &tickers); err != nil {
		return
	}
	return
}

// GetVolumes is used to get the volume for all markets
func (b *Poloniex) GetVolumes() (vc VolumeCollection, err error) {
	r, err := b.client.do("GET", "public?command=return24hVolume", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &vc); err != nil {
		return
	}
	return
}

func (b *Poloniex) GetCurrencies() (currencies Currencies, err error) {
	r, err := b.client.do("GET", "public?command=returnCurrencies", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &currencies.Pair); err != nil {
		return
	}
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC_NXT). 'all' not implemented.
// cat: bid, ask or both to identify the type of orderbook to return.
// depth: how deep of an order book to retrieve
func (b *Poloniex) GetOrderBook(market, cat string, depth int) (orderBook OrderBook, err error) {
	// not implemented
	if cat != "bid" && cat != "ask" && cat != "both" {
		cat = "both"
	}
	if depth > 100 {
		depth = 100
	}
	if depth < 1 {
		depth = 1
	}

	r, err := b.client.do("GET", fmt.Sprintf("public?command=returnOrderBook&currencyPair=%s&depth=%d", strings.ToUpper(market), depth), nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &orderBook); err != nil {
		return
	}
	if orderBook.Error != "" {
		err = errors.New(orderBook.Error)
		return
	}
	return
}

// GetOrderTrades is used to get returns all trades involving a given order
// orderNumber: order number.
func (b *Poloniex) GetOrderTrades(orderNumber int) (tradeOrderTransaction []TradeOrderTransaction, err error) {
	r, err := b.client.doCommand("returnOrderTrades", map[string]string{"orderNumber": fmt.Sprintf("%d", orderNumber)})
	if err != nil {
		return
	}
	if string(r) == `{"error":"Order not found, or you are not the person who placed it."}` {
		err = fmt.Errorf("Error : order not found, or you are not the person who placed it.")
		return
	}
	if err = json.Unmarshal(r, &tradeOrderTransaction); err != nil {
		return
	}
	return
}

// Returns candlestick chart data. Required GET parameters are "currencyPair",
// "period" (candlestick period in seconds; valid values are 300, 900, 1800,
// 7200, 14400, and 86400), "start", and "end". "Start" and "end" are given in
// UNIX timestamp format and used to specify the date range for the data
// returned.
func (b *Poloniex) ChartData(currencyPair string, period int, start, end time.Time) (candles []*CandleStick, err error) {
	r, err := b.client.do("GET", fmt.Sprintf(
		"public?command=returnChartData&currencyPair=%s&period=%d&start=%d&end=%d",
		strings.ToUpper(currencyPair),
		period,
		start.Unix(),
		end.Unix(),
	), nil, false)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &candles); err != nil {
		return
	}

	return
}

func (b *Poloniex) GetBalances() (balances map[string]Balance, err error) {
	balances = make(map[string]Balance)
	r, err := b.client.doCommand("returnCompleteBalances", nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &balances); err != nil {
		return
	}

	return
}

func (b *Poloniex) GetTradeHistory(pair string, start uint32) (trades map[string][]Trade, err error) {
	trades = make(map[string][]Trade)
	r, err := b.client.doCommand("returnTradeHistory", map[string]string{"currencyPair": pair, "start": strconv.FormatUint(uint64(start), 10)})
	if err != nil {
		return
	}

	if pair == "all" {
		if err = json.Unmarshal(r, &trades); err != nil {
			return
		}
	} else {
		var pairTrades []Trade
		if err = json.Unmarshal(r, &pairTrades); err != nil {
			return
		}
		trades[pair] = pairTrades
	}

	return
}

type responseDepositsWithdrawals struct {
	Deposits    []Deposit    `json:"deposits"`
	Withdrawals []Withdrawal `json:"withdrawals"`
}

func (b *Poloniex) GetDepositsWithdrawals(start uint32, end uint32) (deposits []Deposit, withdrawals []Withdrawal, err error) {
	deposits = make([]Deposit, 0)
	withdrawals = make([]Withdrawal, 0)
	r, err := b.client.doCommand("returnDepositsWithdrawals", map[string]string{"start": strconv.FormatUint(uint64(start), 10), "end": strconv.FormatUint(uint64(end), 10)})
	if err != nil {
		return
	}
	var response responseDepositsWithdrawals
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}

	return response.Deposits, response.Withdrawals, nil
}

func (b *Poloniex) Buy(pair string, rate float64, amount float64, tradeType string) (TradeOrder, error) {
	reqParams := map[string]string{
		"currencyPair": pair, "rate": strconv.FormatFloat(rate, 'f', -1, 64),
		"amount": strconv.FormatFloat(amount, 'f', -1, 64)}
	if tradeType != "" {
		reqParams[tradeType] = "1"
	}
	r, err := b.client.doCommand("buy", reqParams)
	if err != nil {
		return TradeOrder{}, err
	}
	var orderResponse TradeOrder
	if err = json.Unmarshal(r, &orderResponse); err != nil {
		return TradeOrder{}, err
	}

	if orderResponse.ErrorMessage != "" {
		return TradeOrder{}, errors.New(orderResponse.ErrorMessage)
	}

	return orderResponse, nil
}

func (b *Poloniex) Sell(pair string, rate float64, amount float64, tradeType string) (TradeOrder, error) {
	reqParams := map[string]string{
		"currencyPair": pair, "rate": strconv.FormatFloat(rate, 'f', -1, 64),
		"amount": strconv.FormatFloat(amount, 'f', -1, 64)}
	if tradeType != "" {
		reqParams[tradeType] = "1"
	}
	r, err := b.client.doCommand("sell", reqParams)
	if err != nil {
		return TradeOrder{}, err
	}
	var orderResponse TradeOrder
	if err = json.Unmarshal(r, &orderResponse); err != nil {
		return TradeOrder{}, err
	}

	if orderResponse.ErrorMessage != "" {
		return TradeOrder{}, errors.New(orderResponse.ErrorMessage)
	}

	return orderResponse, nil
}

func (b *Poloniex) GetOpenOrders(pair string) (openOrders map[string][]OpenOrder, err error) {
	openOrders = make(map[string][]OpenOrder)
	r, err := b.client.doCommand("returnOpenOrders", map[string]string{"currencyPair": pair})
	if err != nil {
		return
	}
	if pair == "all" {
		if err = json.Unmarshal(r, &openOrders); err != nil {
			return
		}
	} else {
		var onePairOrders []OpenOrder
		if err = json.Unmarshal(r, &onePairOrders); err != nil {
			return
		}
		openOrders[pair] = onePairOrders
	}
	return
}

func (b *Poloniex) CancelOrder(orderNumber string) error {
	_, err := b.client.doCommand("cancelOrder", map[string]string{"orderNumber": orderNumber})
	if err != nil {
		return err
	}
	return nil
}

// Returns whole lending history chart data. Required GET parameters are "start",
// "end" (UNIX timestamp format and used to specify the date range for the data returned)
// and optionally limit (<0 for no limit, poloniex automatically limits to 500 records)
func (b *Poloniex) LendingHistory(start, end time.Time, limit int) (lendings []Lending, err error) {
	lendings = make([]Lending, 0)
	reqParams := map[string]string{
		"start": strconv.FormatUint(uint64(start.Unix()), 10),
		"end":   strconv.FormatUint(uint64(end.Unix()), 10)}
	if limit >= 0 {
		reqParams["limit"] = string(limit)
	}

	r, err := b.client.doCommand("returnLendingHistory", reqParams)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &lendings); err != nil {
		return
	}

	return
}
