package poloniex

import (
	"time"
	"encoding/json"
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
