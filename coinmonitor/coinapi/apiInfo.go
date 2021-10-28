package coinapi

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
)

type APIInfo struct {
	rateRemaining int
	rateTotal int
	overageStatus string
}

func (s *Client) ApiInfo() APIInfo{
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/assets", BaseURL), nil) 

	if err != nil {
		panic(err)
	}

	req.Header.Set("X-CoinAPI-Key", s.apiKey)
	req.Header.Del("user-agent")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	reqRemaining, _ := strconv.Atoi(res.Header.Get("X-RateLimit-Remaining")) //Requests remaining able to send to coinapi in 24 hours
	reqTotal, _ := strconv.Atoi(res.Header.Get("X-RateLimit-Limit"))

	log.Printf("req remaining/total 24hr: %v/%v", reqRemaining, reqTotal)
	ret := APIInfo{
		rateRemaining: reqRemaining,
		rateTotal: reqTotal,
		overageStatus: res.Header.Get("X-RateLimit-Overage")}
	return ret
}
