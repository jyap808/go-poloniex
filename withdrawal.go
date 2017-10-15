package poloniex

import (
	"encoding/json"
	"strings"
	"time"
)

type Withdrawal struct {
	WithdrawalNumber uint64    `json:"withdrawalNumber"`
	Currency         string    `json:"currency"`
	Address          string    `json:"address"`
	Amount           float64   `json:"amount,string"`
	Date             time.Time `json:"timestamp"`
	Status           string    `json:"status"`
	TxId             string    `json:"txid"`
	IpAddress        string    `json:"ipAddress"`
}

func (t *Withdrawal) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Withdrawal
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
	if strings.HasPrefix(aux.Status, "COMPLETE") {
		t.TxId = strings.TrimPrefix(t.Status, "COMPLETE: ")
		t.Status = "COMPLETE"
	}
	return nil
}
