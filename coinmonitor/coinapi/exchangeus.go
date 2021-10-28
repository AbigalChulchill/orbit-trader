package coinapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
)

type ExchangeRate struct {
	Time 		 string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate 		 float64 `json:"rate"`
}
  

func (s *Client) USExchangeRate(currency string) float64 {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/exchangerate/USD/%s", BaseURL, currency), nil) 
	
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-CoinAPI-Key", s.apiKey)

	res, err := client.Do(req)
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