package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient        *http.Client
	BaseURL           string
	CollectionsURL    string
	CollectionsApiKey string
}

func (c *Client) GetTxs(address string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address": {address},
		}.Encode()))
}

func (c *Client) GetTxsWithContract(address, contract string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address":  {address},
			"contract": {contract},
		}.Encode()))
}

func (c *Client) getTxs(uri string) (*Page, error) {
	req, _ := http.NewRequest("GET", uri, nil)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		err = errors.E(err, errors.TypePlatformRequest, errors.Params{"uri": uri})
		logger.Error(err, "Ethereum/Trust Ray: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.E("http invalid statuc code", errors.TypePlatformRequest,
			errors.Params{"url": uri, "status_code": res.StatusCode})
	}

	txs := new(Page)
	err = json.NewDecoder(res.Body).Decode(txs)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"uri": uri})
	}
	return txs, nil
}

func (c *Client) GetBlockByNumber(num int64) (page []Doc, err error) {
	path := fmt.Sprintf("%s/transactions/block/%d", c.BaseURL, num)
	res, err := http.Get(path)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"uri": path})
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&page)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"uri": path})
	}

	return page, nil
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	path := fmt.Sprintf("%s/node_info", c.BaseURL)
	res, err := http.Get(path)
	if err != nil {
		return num, errors.E(err, errors.TypePlatformRequest, errors.Params{"uri": path})
	}
	defer res.Body.Close()
	var nodeInfo NodeInfo
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&nodeInfo)
	if err != nil {
		return num, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"uri": path})
	}
	return nodeInfo.LatestBlock, nil
}

func (c *Client) GetTokens(address string) (*TokenPage, error) {
	path := fmt.Sprintf("%s/tokens?%s",
		c.BaseURL,
		url.Values{
			"address": {address},
		}.Encode())

	res, err := http.Get(path)
	if err != nil {
		err = errors.E(err, errors.TypePlatformRequest, errors.Params{"uri": path})
		logger.Error(err, "Ethereum/Trust Ray: Failed to get my tokens")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.E("http invalid statuc code", errors.TypePlatformRequest,
			errors.Params{"url": path, "status_code": res.StatusCode})
	}

	tks := new(TokenPage)
	err = json.NewDecoder(res.Body).Decode(tks)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"uri": path})
	}
	return tks, nil
}
