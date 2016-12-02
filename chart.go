package poloniex

type CandleStick struct {
	Date            PoloniexDate `json:"date"`
	High            float64      `json:"high"`
	Low             float64      `json:"low"`
	Open            float64      `json:"open"`
	Close           float64      `json:"close"`
	Volume          float64      `json:"volume"`
	QuoteVolume     float64      `json:"quoteVolume"`
	WeightedAverage float64      `json:"weightedAverage"`
}
