package main

import (
	"fmt"
	"log"
	"encoding/json"
    "io/ioutil"

	"github.com/CalebCress/orbit-trader/coinmonitor/coinapi"
	krakenapi "github.com/beldur/kraken-go-api-client"
)

type Config struct {
	KRAKEN_API_KEY	  string
	KRAKEN_SECRET_KEY string
	COINAPI_KEY		  string
}

func getConfig() Config {
	content, err := ioutil.ReadFile("./config.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
 
    var config Config
    err = json.Unmarshal(content, &config)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	return config
}

func main() {
	config := getConfig()
	krakenApiKey := config.KRAKEN_API_KEY
	krakenSecretKey := config.KRAKEN_SECRET_KEY
	coinApiKey := config.COINAPI_KEY

	krakenApi := krakenapi.New(krakenApiKey, krakenSecretKey)
	balance, err := krakenApi.Balance(); if err != nil {
		log.Fatal(err)
	}
	api := coinapi.New(coinApiKey)
	fmt.Println(api.USExchangeRate("BTC"))
	log.Printf("balance: %v", balance.USDT)
}
