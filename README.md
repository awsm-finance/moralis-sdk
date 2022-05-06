# Golang Moralis SDK 

## Docs: https://deep-index.moralis.io/api-docs/

-------------

To test requests, create `.env` file, add your api key and run `go run cmd/main.go`


Example:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	moralis "github.com/awsm-finance/moralis-sdk/client"
)

const (
	_host = "https://deep-index.moralis.io/api/v2"
	_address = "0xCC7BcF633f6Ce26cE3eD9E255b8eaA6f219A0956"
)

func main() {
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
}
```