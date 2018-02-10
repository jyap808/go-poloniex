package poloniex

type Tickers struct {
	Pair map[string]Ticker
}

type Ticker struct {
	Id            int     `json:"id"`
	Last          float64 `json:"last,string"`
	LowestAsk     float64 `json:"lowestAsk,string"`
	HighestBid    float64 `json:"highestBid,string"`
	PercentChange float64 `json:"percentChange,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
	IsFrozen      int     `json:"isFrozen,string"`
	High24Hr      float64 `json:"high24hr,string"`
	Low24Hr       float64 `json:"low24hr,string"`
}
