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

	respErc20Balance, err := c.GetERC20BalanceByAddress(&moralis.GetERC20BalanceByAddressInput{
		Address:        _address,
		Chain:          moralis.ChainEth,
		TokenAddresses: []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"},
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, balance := range respErc20Balance {
		fmt.Printf("%d. GetERC20BalanceByAddress (%s): %+v\n", i+1, _address, balance)
	}

	respErc20Transfers, err := c.GetERC20TransfersByAddress(&moralis.GetERC20TransfersByAddressInput{
		Address: _address,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GetERC20TransfersByAddress (%s): %+v\n", _address, respErc20Transfers)
}
