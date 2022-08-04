package moralisapi

import (
	"errors"
	"net/url"
	"strconv"
)

type GetTransactionsByAddressInput struct {
	Chain     string
	Address   string
	FromBlock int
	ToBlock   int
	FromDate  string
	ToDate    string
	Offset    int
	Cursor    string
	Limit     int
}

func (r GetTransactionsByAddressInput) Validate() error {
	if r.Address == "" {
		return errors.New("missing address in GetTransactionsByAddressInput")
	}

	return nil
}

func (r GetTransactionsByAddressInput) Query() string {
	u := url.URL{}
	values := u.Query()

	if r.Chain != "" {
		values.Add("chain", r.Chain)
	}

	if r.FromBlock != 0 {
		values.Add("from_block", strconv.Itoa(r.FromBlock))
	}

	if r.ToBlock != 0 {
		values.Add("to_block", strconv.Itoa(r.ToBlock))
	}

	if r.FromDate != "" {
		values.Add("from_date", r.FromDate)
	}

	if r.ToDate != "" {
		values.Add("to_date", r.ToDate)
	}

	if r.Offset != 0 {
		values.Add("offset", strconv.Itoa(r.Offset))
	}

	if r.Cursor != "" {
		values.Add("cursor", r.Cursor)
	}

	if r.Limit != 0 {
		values.Add("limit", strconv.Itoa(r.Limit))
	}

	return values.Encode()
}

type GetTransactionsByAddressResponse struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Result   []Transaction `json:"result"`
}

type Transaction struct {
	Hash           string `json:"hash"`
	FromAddress    string `json:"from_address"`
	ToAddress      string `json:"to_address"`
	Value          string `json:"value"`
	Gas            string `json:"gas"`
	GasPrice       string `json:"gas_price"`
	BlockTimestamp string `json:"block_timestamp"`
	BlockNumber    string `json:"block_number"`
	BlockHash      string `json:"block_hash"`
}

type GetBalanceByAddressInput struct {
	Chain   string
	Address string
	ToBlock int
}

func (r GetBalanceByAddressInput) Validate() error {
	if r.Address == "" {
		return errors.New("missing address in GetTransactionsByAddressInput")
	}

	return nil
}

func (r GetBalanceByAddressInput) Query() string {
	u := url.URL{}
	values := u.Query()

	if r.Chain != "" {
		values.Add("chain", r.Chain)
	}

	if r.ToBlock != 0 {
		values.Add("to_block", strconv.Itoa(r.ToBlock))
	}

	return values.Encode()
}

type GetBalanceByAddressResponse struct {
	Balance string `json:"balance"`
}

type GetERC20BalanceByAddressInput struct {
	Chain          string
	Address        string
	ToBlock        int
	TokenAddresses []string
}

func (r GetERC20BalanceByAddressInput) Validate() error {
	if r.Address == "" {
		return errors.New("missing address in GetTransactionsByAddressInput")
	}

	return nil
}

func (r GetERC20BalanceByAddressInput) Query() string {
	u := url.URL{}
	values := u.Query()

	if r.Chain != "" {
		values.Add("chain", r.Chain)
	}

	if r.ToBlock != 0 {
		values.Add("to_block", strconv.Itoa(r.ToBlock))
	}

	if r.TokenAddresses != nil && len(r.TokenAddresses) > 0 {
		for _, address := range r.TokenAddresses {
			values.Add("token_addresses", address)
		}
	}

	return values.Encode()
}

type GetERC20BalanceByAddressResponse struct {
	TokenAddress string `json:"token_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Logo         string `json:"logo"`
	Thumbnail    string `json:"thumbnail"`
	Decimals     int    `json:"decimals"`
	Balance      string `json:"balance"`
}

type GetERC20TransfersByAddressInput struct {
	Chain     string
	Address   string
	FromBlock int
	ToBlock   int
	FromDate  string
	ToDate    string
	Offset    int
	Limit     int
	Cursor    string
}

func (r GetERC20TransfersByAddressInput) Validate() error {
	if r.Address == "" {
		return errors.New("missing address in GetTransactionsByAddressInput")
	}

	return nil
}

func (r GetERC20TransfersByAddressInput) Query() string {
	u := url.URL{}
	values := u.Query()

	if r.Chain != "" {
		values.Add("chain", r.Chain)
	}

	if r.FromBlock != 0 {
		values.Add("from_block", strconv.Itoa(r.FromBlock))
	}

	if r.ToBlock != 0 {
		values.Add("to_block", strconv.Itoa(r.ToBlock))
	}

	if r.FromDate != "" {
		values.Add("from_date", r.FromDate)
	}

	if r.ToDate != "" {
		values.Add("to_date", r.ToDate)
	}

	if r.Offset != 0 {
		values.Add("offset", strconv.Itoa(r.Offset))
	}

	if r.Limit != 0 {
		values.Add("limit", strconv.Itoa(r.Limit))
	}

	return values.Encode()
}

type GetERC20TransfersByAddressResponse struct {
	Total    int                  `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
	Result   []ERC20TokenTransfer `json:"result"`
}

type ERC20TokenTransfer struct {
	TransactionHash string `json:"transaction_hash"`
	Address         string `json:"address"` // address of a ERC20 token smart contract
	BlockTimestamp  string `json:"block_timestamp"`
	BlockNumber     string `json:"block_number"`
	BlockHash       string `json:"block_hash"`
	ToAddress       string `json:"to_address"`
	FromAddress     string `json:"from_address"`
	Value           string `json:"value"`
}

type GetTransactionByHashInput struct {
	Chain string
	Hash  string
}

func (r GetTransactionByHashInput) Validate() error {
	if r.Hash == "" {
		return errors.New("missing hash in GetTransactionByHashInput")
	}

	return nil
}

func (r GetTransactionByHashInput) Query() string {
	u := url.URL{}
	values := u.Query()

	if r.Chain != "" {
		values.Add("chain", r.Chain)
	}
	return values.Encode()
}
