package iotex

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
)

type Client struct {
	blockatlas.Request
}

func InitClient(baseUrl string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
			BaseUrl:      baseUrl,
		},
	}
}

func (c *Client) GetLatestBlock() (int64, error) {
	var chainMeta ChainMeta
	err := c.Get(&chainMeta, "chainmeta", nil)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(chainMeta.Height, 10, 64)
}

func (c *Client) GetTxsInBlock(number int64) ([]*ActionInfo, error) {
	path := fmt.Sprintf("transfers/block/%d", number)
	var resp Response
	err := c.Get(&resp, path, nil)
	if err != nil {
		return nil, err
	}
	return resp.ActionInfo, nil
}

func (c *Client) GetTxsOfAddress(address string, start int64) (*Response, error) {
	var response Response
	err := c.Get(&response, "actions/addr/"+address, url.Values{
		"start": {strconv.FormatInt(start, 10)},
		"count": {strconv.FormatInt(blockatlas.TxPerPage, 10)},
	})

	if err != nil {
		logger.Error(err, "IOTEX: Failed to get transactions for address", logger.Params{"address": address})
		return nil, blockatlas.ErrSourceConn
	}
	return &response, err
}

func (c *Client) GetAddressTotalTransactions(address string) (int64, error) {
	var account AccountInfo
	err := c.Get(&account, "accounts/"+address, nil)
	if err != nil {
		return 0, nil
	}
	numActions, err := strconv.ParseInt(account.AccountMeta.NumActions, 10, 64)
	if err != nil {
		return 0, nil
	}

	return numActions, nil
}
