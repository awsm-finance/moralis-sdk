package moralis

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
