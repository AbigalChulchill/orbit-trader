package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CalebCress/orbit-trader/coinmonitor/coinapi"
	krakenapi "github.com/beldur/kraken-go-api-client"
)

func getEnv(key string) string{
	val, ok := os.LookupEnv(key)
	if !ok {
		v := key
		panic(fmt.Errorf("%v not set\n", v))
	} else {
		return(val)
	}
}

func main() {
	krakenApiKey := getEnv("KRAKEN_API_KEY")
	krakenSecretKey := getEnv("KRAKEN_SECRET_KEY")
	coinApiKey := getEnv("COINAPI_KEY")

	krakenApi := krakenapi.New(krakenApiKey, krakenSecretKey)
	balance, err := krakenApi.Balance(); if err != nil {
		log.Fatal(err)
	}
	api := coinapi.New(coinApiKey)
	fmt.Println(api.usExchangeRate("BTC"))
	log.Printf("balance: %v", balance.USDT)
}
