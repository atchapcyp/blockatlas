package tron

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas"
)

type Page struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Txs     []Tx   `json:"data"`
}

type Tx struct {
	ID        string `json:"txID"`
	BlockTime int64  `json:"block_timestamp"`
	Data      TxData `json:"raw_data"`
}

type TxData struct {
	Contracts []Contract `json:"contract"`
}

type Contract struct {
	Type      string      `json:"type"`
	Parameter interface{} `json:"parameter"`
}

type TransferContract struct {
	Value TransferValue `json:"value"`
}

type TransferValue struct {
	Amount       blockatlas.Amount `json:"amount"`
	OwnerAddress string            `json:"owner_address"`
	ToAddress    string            `json:"to_address"`
}

type Accounts struct {
	Data []AccountsData `json:data`
}

type AccountsData struct {
	AssetsV2 []AssetV2 `json:"assetV2"`
}

type AssetV2 struct {
	Key string `json:"key"`
}

type Asset struct {
	Data []AssetInfo `json:"data"`
}

type AssetInfo struct {
	Name     string `json:"name"`
	Symbol   string `json:"abbr"`
	ID       string `json:"id"`
	Decimals uint   `json:"precision"`
}

type Validators struct {
	Witnesses []Validator `json:"witnesses"`
}

type Validator struct {
	Address string `json:"address"`
}

func (c *Contract) UnmarshalJSON(buf []byte) error {
	var contractInternal struct {
		Type      string          `json:"type"`
		Parameter json.RawMessage `json:"parameter"`
	}
	err := json.Unmarshal(buf, &contractInternal)
	if err != nil {
		return err
	}
	switch contractInternal.Type {
	case "TransferContract":
		var transfer TransferContract
		err = json.Unmarshal(contractInternal.Parameter, &transfer)
		c.Parameter = transfer
	}
	return err
}
