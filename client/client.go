package moralis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	_getTransactionsByAddress = "/%s"         // /{address}
	_getBalanceByAddress      = "/%s/balance" // /{address}/balance
)

type Client struct {
	host   string
	apiKey string

	httpClient *http.Client
}

func NewClient(host, apiKey string, timeout time.Duration) *Client {
	return &Client{
		host, apiKey, &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetTransactionsByAddress(inp *GetTransactionsByAddressInput) (*GetTransactionsByAddressResponse, error) {
	if err := inp.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(_getTransactionsByAddress, inp.Address)

	if q := inp.Query(); q != "" {
		path = fmt.Sprintf("%s?%s", path, q)
	}

	var resp GetTransactionsByAddressResponse
	err := c.request(path, http.MethodGet, &resp)

	return &resp, err
}

func (c *Client) GetBalanceByAddress(inp *GetBalanceByAddressInput) (*GetBalanceByAddressResponse, error) {
	if err := inp.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(_getBalanceByAddress, inp.Address)

	if q := inp.Query(); q != "" {
		path = fmt.Sprintf("%s?%s", path, q)
	}

	var resp GetBalanceByAddressResponse
	err := c.request(path, http.MethodGet, &resp)

	return &resp, err
}

func (c *Client) request(path, method string, out interface{}) error {
	req, err := http.NewRequest(method, c.host+path, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("non-2XX response")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, out)
}
