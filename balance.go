package poloniex

type Balance struct {
	Available string `json:"available"`
	BtcValue  string `json:"btcValue"`
	OnOrders  string `json:"onOrders"`
}
