package coinapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type ExchangeRate struct {
	Time 		 string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate 		 float64 `json:"rate"`
}
  

func (s *Client) usExchangeRate(currency string) float64 {
	res, err := http.Get(fmt.Sprintf("%s/v1/exchangerate/USD/%s", BaseURL, s.apiKey)) 
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var decoded ExchangeRate
	if err := json.NewDecoder(res.Body).Decode(&decoded); err != nil {
		panic(err)
	}
	return decoded.Rate
}