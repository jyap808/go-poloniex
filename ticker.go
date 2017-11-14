package poloniex

import "github.com/shopspring/decimal"

type Tickers struct {
	Pair map[string]Ticker
}

type Ticker struct {
	Last          decimal.Decimal `json:"last,string"`
	LowestAsk     decimal.Decimal `json:"lowestAsk,string"`
	HighestBid    decimal.Decimal `json:"highestBid,string"`
	PercentChange decimal.Decimal `json:"percentChange,string"`
	BaseVolume    decimal.Decimal `json:"baseVolume,string"`
	QuoteVolume   decimal.Decimal `json:"quoteVolume,string"`
	IsFrozen      int             `json:"isFrozen,string"`
	High24Hr      decimal.Decimal `json:"high24hr,string"`
	Low24Hr       decimal.Decimal `json:"low24hr,string"`
}
