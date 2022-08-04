package moralisapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	_getTransactionsByAddress   = "/%s"                 // /{address}
	_getBalanceByAddress        = "/%s/balance"         // /{address}/balance
	_getERC20BalanceByAddress   = "/%s/erc20"           // /{address}/erc20
	_getERC20TransfersByAddress = "/%s/erc20/transfers" // /{address}/erc20/transfers
	_getTransactionByHash       = "/transaction/%s"     // /transaction/{hash}
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

func (c *Client) GetERC20BalanceByAddress(inp *GetERC20BalanceByAddressInput) ([]*GetERC20BalanceByAddressResponse, error) {
	if err := inp.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(_getERC20BalanceByAddress, inp.Address)

	if q := inp.Query(); q != "" {
		path = fmt.Sprintf("%s?%s", path, q)
	}

	resp := make([]*GetERC20BalanceByAddressResponse, 0)
	err := c.request(path, http.MethodGet, &resp)

	return resp, err
}

func (c *Client) GetERC20TransfersByAddress(inp *GetERC20TransfersByAddressInput) (*GetERC20TransfersByAddressResponse, error) {
	if err := inp.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(_getERC20TransfersByAddress, inp.Address)

	if q := inp.Query(); q != "" {
		path = fmt.Sprintf("%s?%s", path, q)
	}

	var resp GetERC20TransfersByAddressResponse
	err := c.request(path, http.MethodGet, &resp)

	return &resp, err
}

func (c *Client) GetTransactionByHash(inp *GetTransactionByHashInput) (*Transaction, error) {
	if err := inp.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(_getTransactionByHash, inp.Hash)

	if q := inp.Query(); q != "" {
		path = fmt.Sprintf("%s?%s", path, q)
	}

	var resp Transaction
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
		return fmt.Errorf("request failed, status code %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, out)
}
