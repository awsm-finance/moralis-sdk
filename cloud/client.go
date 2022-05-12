package moraliscloud

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

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

func (c *Client) WatchEthAddress(address string) error {
	q := urlQuery(address)
	return c.request(_watchEthAddressUrl, http.MethodGet, q)
}

func (c *Client) UnwatchEthAddress(address string) error {
	q := urlQuery(address)
	return c.request(_unwatchEthAddressUrl, http.MethodGet, q)
}

func (c *Client) WatchBscAddress(address string) error {
	q := urlQuery(address)
	return c.request(_watchBscAddressUrl, http.MethodGet, q)
}

func (c *Client) UnwatchBscAddress(address string) error {
	q := urlQuery(address)
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

func urlQuery(address string) url.Values {
	u := url.URL{}

	query := u.Query()
	query.Add("address", address)

	return query
}
