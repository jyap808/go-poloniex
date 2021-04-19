package poloniex

import (
	"encoding/json"
	"strconv"
)

type Volume map[string]float64

type VolumeCollection struct {
	TotalBNB  float64 `json:"totalBNB,string"`
	TotalBTC  float64 `json:"totalBTC,string"`
	TotalBUSD float64 `json:"totalBUSD,string"`
	TotalDAI  float64 `json:"totalDAI,string"`
	TotalETH  float64 `json:"totalETH,string"`
	TotalPAX  float64 `json:"totalPAX,string"`
	TotalTRX  float64 `json:"totalTRX,string"`
	TotalTUSD float64 `json:"totalTUSD,string"`
	TotalUSDC float64 `json:"totalUSDC,string"`
	TotalUSDJ float64 `json:"totalUSDJ,string"`
	TotalUSDT float64 `json:"totalUSDT,string"`
	TotalXMR  float64 `json:"totalXMR,string"`
	TotalXUSD float64 `json:"totalXUSD,string"`
	Volumes   map[string]Volume
}

func (tc *VolumeCollection) UnmarshalJSON(b []byte) error {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	tc.Volumes = make(map[string]Volume)
	for k, v := range m {
		switch k {
		case "totalBNB":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalBNB = f
		case "totalBTC":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalBTC = f
		case "totalBUSD":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalBUSD = f
		case "totalDAI":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalDAI = f
		case "totalETH":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalETH = f
		case "totalPAX":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalPAX = f
		case "totalTRX":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalTRX = f
		case "totalTUSD":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalTUSD = f
		case "totalUSDC":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalUSDC = f
		case "totalUSDJ":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalUSDJ = f
		case "totalUSDT":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalUSDT = f
		case "totalXMR":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalXMR = f
		case "totalXUSD":
			f, err := parseJSONFloatString(v)
			if err != nil {
				return err
			}
			tc.TotalXUSD = f
		default:
			t := make(Volume)
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			tc.Volumes[k] = t
		}
	}
	return nil
}

func (t *Volume) UnmarshalJSON(b []byte) error {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	for k, v := range m {
		f, err := parseJSONFloatString(v)
		if err != nil {
			return err
		}
		(*t)[k] = f
	}
	return nil
}

func parseJSONFloatString(b json.RawMessage) (float64, error) {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(s, 64)
}
