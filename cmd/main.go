package main

import (
	"fmt"
	"log"
	"os"
	"time"

	moraliscloud "github.com/awsm-finance/moralis-sdk/cloud"
	moralisapi "github.com/awsm-finance/moralis-sdk/restapi"

	"github.com/joho/godotenv"
)

const (
	_host       = "https://deep-index.moralis.io/api/v2"
	_serverHost = "https://wamaxhbnkkbj.usemoralis.com:2053/server"
	// _serverHost = "https://mwq7lmluczgx.usemoralis.com:2053/server"

	_address = "0xcc7bcf633f6ce26ce3ed9e255b8eaa6f219a0956"
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

	appId := os.Getenv("CLOUD_APP_ID")
	if appId == "" {
		log.Fatal("CLOUD_APP_ID is empty")
	}

	masterKey := os.Getenv("CLOUD_MASTER_KEY")
	if appId == "" {
		log.Fatal("CLOUD_MASTER_KEY is empty")
	}

	c := moralisapi.NewClient(_host, apiKey, time.Second*5)

	tx, err := c.GetTransactionByHash(&moralisapi.GetTransactionByHashInput{
		Hash: "0x25cd7f3bfc73aa0fa3ea1307b0271cb78d703594389462f596fc033ca5170c2b",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GetTransactionByHash: %+v", tx)

	// resp, err := c.GetTransactionsByAddress(&moralisapi.GetTransactionsByAddressInput{
	// 	Address: _address,
	// 	Chain:   moralisapi.ChainRopsten,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetTransactionsByAddress (%s): %+v\n", _address, resp)

	// respBalance, err := c.GetBalanceByAddress(&moralisapi.GetBalanceByAddressInput{
	// 	Address: _address,
	// 	Chain:   moralisapi.ChainEth,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetBalanceByAddress (%s): %+v\n", _address, respBalance)

	respErc20Balance, err := c.GetERC20BalanceByAddress(&moralisapi.GetERC20BalanceByAddressInput{
		Address: _address,
		Chain:   moralisapi.ChainEth,
		// TokenAddresses: []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"},
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, balance := range respErc20Balance {
		fmt.Printf("%d. GetERC20BalanceByAddress (%s): %+v\n", i+1, _address, balance)
	}

	// respErc20Transfers, err := c.GetERC20TransfersByAddress(&moralisapi.GetERC20TransfersByAddressInput{
	// 	Address: _address,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetERC20TransfersByAddress (%s): %+v\n", _address, respErc20Transfers)

	cloudClient := moraliscloud.NewClient(_serverHost, appId, masterKey, time.Second*5)
	if err := cloudClient.WatchEthAddress(_address, true); err != nil {
		log.Printf("error: %s", err.Error())
	} else {
		log.Println("address registered!!!!!!!")
	}
}
