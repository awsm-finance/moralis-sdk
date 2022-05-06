package main

import (
	"fmt"
	"log"
	"os"
	"time"

	moralis "github.com/awsm-finance/moralis-sdk/client"
	"github.com/joho/godotenv"
)

const (
	_host = "https://deep-index.moralis.io/api/v2"

	_address = "0xCC7BcF633f6Ce26cE3eD9E255b8eaA6f219A0956"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is empty")
	}

	c := moralis.NewClient(_host, apiKey, time.Second*5)

	resp, err := c.GetTransactionsByAddress(&moralis.GetTransactionsByAddressInput{
		Address: _address,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GetTransactionsByAddress (%s): %+v\n", _address, resp)

	respBalance, err := c.GetBalanceByAddress(&moralis.GetBalanceByAddressInput{
		Address: _address,
		Chain:   moralis.ChainEth,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GetBalanceByAddress (%s): %+v\n", _address, respBalance)
}
