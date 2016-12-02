package poloniex

import (
	"encoding/binary"
	"errors"
	"time"
)

type PoloniexDate struct {
	time.Time
}

func (pd *PoloniexDate) UnmarshalJSON(data []byte) error {
	i, err := binary.Varint(data)
	if err == 0 {
		return errors.New("Timestamp invalid (can't parse int64)")
	}
	pd.Time = time.Unix(i, 0)
	return nil
}
