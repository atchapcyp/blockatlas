package tron

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

type Platform struct {
	client Client
}

const Annual = 4.32

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("tron.api"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range srcTxs {
		tx, ok := Normalize(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()

	if err != nil {
		return results, err
	}

	for _, v := range validators.Witnesses {
		if val, ok := normalizeValidator(v); ok {
			results = append(results, val)
		}
	}

	return results, nil
}

func normalizeValidator(v Validator) (validator blockatlas.Validator, ok bool) {
	address, err := HexToAddress(v.Address)
	if err != nil {
		return validator, false
	}

	return blockatlas.Validator{
		Status: true,
		ID:     address,
		Reward: blockatlas.StakingReward{Annual: Annual},
	}, true
}

/// Normalize converts a Tron transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	if len(srcTx.Data.Contracts) < 1 {
		return tx, false
	}

	// TODO Support multiple transfers in a single transaction
	contract := &srcTx.Data.Contracts[0]
	switch contract.Parameter.(type) {
	case TransferContract:
		transfer := contract.Parameter.(TransferContract)
		from, err := HexToAddress(transfer.Value.OwnerAddress)
		if err != nil {
			return tx, false
		}
		to, err := HexToAddress(transfer.Value.ToAddress)
		if err != nil {
			return tx, false
		}

		return blockatlas.Tx{
			ID:   srcTx.ID,
			Coin: coin.TRX,
			Date: srcTx.BlockTime / 1000,
			From: from,
			To:   to,
			Fee:  "0",
			Meta: blockatlas.Transfer{
				Value:    transfer.Value.Amount,
				Symbol:   coin.Coins[coin.TRX].Symbol,
				Decimals: coin.Coins[coin.TRX].Decimals,
			},
		}, true
	default:
		return tx, false
	}
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	tokens, err := p.client.GetAccountMetadata(address)
	if err != nil {
		return nil, err
	}

	tokenPage := make([]blockatlas.Token, 0)
	var tokenIDs []string
	if len(tokens.Data) > 0 {
		for _, v := range tokens.Data[0].AssetsV2 {
			tokenIDs = append(tokenIDs, v.Key)
		}
	} else {
		return tokenPage, nil
	}

	tokensInfoChan := make(chan *Asset, len(tokenIDs))

	var wg sync.WaitGroup
	wg.Add(len(tokenIDs))
	for _, id := range tokenIDs {
		go func(id string) {
			defer wg.Done()
			info, err := p.client.GetTokenInfo(id)
			if err != nil {
				logger.Error("GetTokenInfo", err)
			}
			tokensInfoChan <- info
		}(id)
	}
	wg.Wait()
	close(tokensInfoChan)

	tokensInfoMap := make(map[string]AssetInfo)
	for info := range tokensInfoChan {
		tokensInfoMap[info.Data[0].ID] = info.Data[0]
	}

	for _, v := range tokens.Data[0].AssetsV2 {
		tokenPage = append(tokenPage, NormalizeToken(tokensInfoMap[v.Key]))
	}

	return tokenPage, nil
}

func NormalizeToken(info AssetInfo) blockatlas.Token {
	return blockatlas.Token{
		Name:     info.Name,
		Symbol:   info.Symbol,
		TokenID:  info.ID,
		Coin:     coin.TRX,
		Decimals: info.Decimals,
		Type:     blockatlas.TokenTypeTRC10,
	}
}
