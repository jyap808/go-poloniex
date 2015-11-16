package poloniex

type Currency struct {
	Name               string  `json:"name"`
	MaxDailyWithdrawal float64 `json:"maxDailyWithdrawal,string"`
	TxFee              float64 `json:"txFee,string"`
	MinConf            int     `json:"minConf"`
	Disabled           int     `json:"disabled"`
	Frozen             int     `json:"frozen"`
	Delisted           int     `json:"delisted"`
}

type Currencies struct {
	Pair map[string]Currency
}
