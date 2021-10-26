package coinmonitor

import (
	krakenapi "github.com/beldur/kraken-go-api-client"
)

type Monitor struct {
	coin string
	krakenapi *krakenapi.KrakenAPI
}

func New(coin string, krakenapi *krakenapi.KrakenAPI) *Monitor {
	return &Monitor{coin: coin, krakenapi: krakenapi}
}