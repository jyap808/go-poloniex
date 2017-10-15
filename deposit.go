package poloniex

import (
	"encoding/json"
	"time"
)

type Deposit struct {
	Currency      string    `json:"currency"`
	Address       string    `json:"address"`
	Amount        float64   `json:"amount,string"`
	Confirmations uint64    `json:"confirmations"`
	TxId          string    `json:"txid"`
	Date          time.Time `json:"timestamp"`
	Status        string    `json:"status"`
}

func (t *Deposit) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Deposit
	aux := &struct {
		Date int64 `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Date = time.Unix(aux.Date, 0)
	return nil
}
