package blockatlas

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/coin"
)

// Initer is a service that can be initialized once
type Initer interface {
	Init() error
}

// Platform can be used to access a crypto service
type Platform interface {
	Initer
	Coin() coin.Coin
}

// TxAPI provides transaction lookups
type TxAPI interface {
	Platform
	GetTxsByAddress(address string) (TxPage, error)
}

// TokenTxAPI provides token transaction lookups
type TokenTxAPI interface {
	Platform
	GetTokenTxsByAddress(address string, token string) (TxPage, error)
}

// TokenAPI provides token lookups
type TokenAPI interface {
	Platform
	GetTokenListByAddress(address string) (TokenPage, error)
}

// BlockAPI provides block information and lookups
type BlockAPI interface {
	Platform
	CurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64) (*Block, error)
}

// StakingAPI provides staking information
type StakeAPI interface {
	Platform
	GetValidators() (ValidatorPage, error)
}

type CollectionAPI interface {
	Platform
	GetCollections(owner string) (CollectionPage, error)
	GetCollectibles(owner, collectibleID string) (CollectiblePage, error)
}

// CustomAPI provides custom HTTP routes
type CustomAPI interface {
	Platform
	RegisterRoutes(router gin.IRouter)
}
