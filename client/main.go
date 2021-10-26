package main

import (
	krakenapi "github.com/beldur/kraken-go-api-client"
	"os"
	"fmt"
	"log"
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

	api := krakenapi.New(krakenApiKey, krakenSecretKey)
	balance, err := api.Balance(); if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("balance: %v", balance.USDT)
}
