package moraliscloud

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Docs: https://docs.moralis.io/moralis-dapp/automatic-transaction-sync/historical-transactions

const (
	_watchEthAddressUrl   = "/functions/watchEthAddress"
	_unwatchEthAddressUrl = "/functions/unwatchEthAddress"
	_watchBscAddressUrl   = "/functions/watchBscAddress"
	_unwatchBscAddressUrl = "/functions/unwatchBscAddress"
)

type Client struct {
	serverHost string // e.g. https://wamaxhbnkkbj.usemoralis.com:2053/server
	appId      string
	masterKey  string

	httpClient *http.Client
}

func NewClient(serverHost, appId, masterKey string, timeout time.Duration) *Client {
	return &Client{serverHost, appId, masterKey, &http.Client{Timeout: timeout}}
}

func (c *Client) WatchEthAddress(address string, syncHistory bool) error {
	q := urlQuery(address, syncHistory)
	return c.request(_watchEthAddressUrl, http.MethodGet, q)
}

func (c *Client) UnwatchEthAddress(address string) error {
	q := urlQuery(address, false)
	return c.request(_unwatchEthAddressUrl, http.MethodGet, q)
}

func (c *Client) WatchBscAddress(address string, syncHistory bool) error {
	q := urlQuery(address, syncHistory)
	return c.request(_watchBscAddressUrl, http.MethodGet, q)
}

func (c *Client) UnwatchBscAddress(address string) error {
	q := urlQuery(address, false)
	return c.request(_unwatchBscAddressUrl, http.MethodGet, q)
}

func (c *Client) request(path, method string, queryParams url.Values) error {
	queryParams.Add("_ApplicationId", c.appId)

	url := fmt.Sprintf("%s%s?%s", c.serverHost, path, queryParams.Encode())

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Parse-Master-Key", c.masterKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("request failed, status code %d", resp.StatusCode)
	}

	return nil
}

func urlQuery(address string, syncHistory bool) url.Values {
	u := url.URL{}

	query := u.Query()
	query.Add("address", address)

	// it's "true" by default
	if !syncHistory {
		query.Add("Sync_historical", "false")
	}

	return query
}
