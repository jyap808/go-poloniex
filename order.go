package poloniex

type OrderBook struct {
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
	IsFrozen int             `json:"isFrozen,string"`
	Error    string          `json:"error"`
}

// This can probably be implemented using UnmarshalJSON
/*
type OrderBook struct {
	Bids     []Orderb `json:"bids"`
	Asks     []Orderb `json:"asks"`
	IsFrozen int      `json:"isFrozen,string"`
}
type Orderb struct {
	Rate     string
	Quantity float64
}
*/
