package poloniex

import (
	"encoding/json"
	"time"
)

type Lending struct {
	Id       uint64    `json:"id"`
	Currency string    `json:"currency"`
	Rate     float64   `json:"rate,string"`
	Amount   float64   `json:"amount,string"`
	Duration float64   `json:"duration,string"`
	Interest float64   `json:"interest,string"`
	Fee      float64   `json:"fee,string"`
	Earned   float64   `json:"earned,string"`
	Open     time.Time `json:"open,string"`
	Close    time.Time `json:"close,string"`
}

func (t *Lending) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Lending
	aux := &struct {
		Open  string `json:"open"`
		Close string `json:"close"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Open, err = time.Parse("2006-01-02 15:04:05", aux.Open)
	t.Close, err = time.Parse("2006-01-02 15:04:05", aux.Close)
	if err != nil {
		return err
	}
	return nil
}
