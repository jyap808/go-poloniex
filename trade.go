package poloniex

import (
	"encoding/json"
	"time"
)

const (
	TRADE_FILL_OR_KILL        = "fillOrKill"
	TRADE_IMMEDIATE_OR_CANCEL = "immediateOrCancel"
	TRADE_POST_ONLY           = "postOnly"
)

type Trade struct {
	GlobalTradeID uint64    `json:"globalTradeID"`
	TradeID       uint64    `json:"tradeID,string"`
	Date          time.Time `json:"date,string"`
	Type          string    `json:"type"`
	Category      string    `json:"category"`
	Rate          float64   `json:"rate,string"`
	Amount        float64   `json:"amount,string"`
	Total         float64   `json:"total,string"`
	Fee           float64   `json:"fee,string"`
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Trade
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Date, err = time.Parse("2006-01-02 15:04:05", aux.Date)
	if err != nil {
		return err
	}
	return nil
}

type ResultingTrade struct {
	Amount  float64 `json:"amount,string"`
	Date    string  `json:"date"`
	Rate    float64 `json:"rate,string"`
	Total   float64 `json:"total,string"`
	TradeID string  `json:"tradeID"`
	Type    string  `json:"type"`
}

type TradeOrder struct {
	OrderNumber     string           `json:"orderNumber"`
	ResultingTrades []ResultingTrade `json:"resultingTrades"`
	ErrorMessage    string           `json:"error"`
}

type TradeOrderTransaction struct {
	GlobalTradeID uint64    `json:"globalTradeID"`
	TradeID       uint64    `json:"tradeID"`
	CurrencyPair  string    `json:"currencyPair"`
	Type          string    `json:"type"`
	Rate          float64   `json:"rate,string"`
	Amount        float64   `json:"amount,string"`
	Total         float64   `json:"total,string"`
	Fee           float64   `json:"fee,string"`
	Date          time.Time `json:"date,string"`
}
