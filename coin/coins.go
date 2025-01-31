// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-09-20 22:51:29.143552 -0700 PDT m=&#43;0.001381525
// using data from coins.yml
package coin

import (
	"fmt"
)

// Coin is the native currency of a blockchain
type Coin struct {
	ID               uint
	Handle           string
	Symbol           string
	Name             string
	Decimals         uint
	BlockTime        int
	MinConfirmations int64
	SampleAddr       string
}

func (c *Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Name, c.ID)
}

const (
	ETH = 60
	ETC = 61
	ICX = 74
	ATOM = 118
	XRP = 144
	XLM = 148
	POA = 178
	TRX = 195
	FIO = 235
	NIM = 242
	IOTX = 304
	ZIL = 313
	AION = 425
	AE = 457
	THETA = 500
	BNB = 714
	VET = 818
	CLO = 820
	TOMO = 889
	TT = 1001
	ONT = 1024
	XTZ = 1729
	KIN = 2017
	NAS = 2718
	GO = 6060
	WAN = 5718350
	WAVES = 5741564
	SEM = 7562605
	BTC = 0
	LTC = 2
	DOGE = 3
	DASH = 5
	VIA = 14
	GRS = 17
	ZEC = 133
	XZC = 136
	BCH = 145
	RVN = 175
	QTUM = 2301
	ZEL = 19167
	DCR = 42
)

var Coins = map[uint]Coin{
	ETH: {
		ID:               60,
		Handle:           "ethereum",
		Symbol:           "ETH",
		Name:             "Ethereum",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1",
	},
	ETC: {
		ID:               61,
		Handle:           "classic",
		Symbol:           "ETC",
		Name:             "Ethereum Classic",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "0xf3524415b6D873205B4c3Cda783527b2aC4daAA9",
	},
	ICX: {
		ID:               74,
		Handle:           "icon",
		Symbol:           "ICX",
		Name:             "ICON",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
	},
	ATOM: {
		ID:               118,
		Handle:           "cosmos",
		Symbol:           "ATOM",
		Name:             "Cosmos ATOM",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
	},
	XRP: {
		ID:               144,
		Handle:           "ripple",
		Symbol:           "XRP",
		Name:             "Ripple",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
	},
	XLM: {
		ID:               148,
		Handle:           "stellar",
		Symbol:           "XLM",
		Name:             "Stellar Lumens",
		Decimals:         7,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	},
	POA: {
		ID:               178,
		Handle:           "poa",
		Symbol:           "POA",
		Name:             "Poa",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "0x1fddEc96688e0538A316C64dcFd211c491ECf0d8",
	},
	TRX: {
		ID:               195,
		Handle:           "tron",
		Symbol:           "TRX",
		Name:             "Tron",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	},
	FIO: {
		ID:               235,
		Handle:           "fio",
		Symbol:           "FIO",
		Name:             "FIO",
		Decimals:         9,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "FIO5J2xdfWygeNdHZNZRzRws8YGbVxjUXtp4eP8KoGkGKoLFQ7CaU",
	},
	NIM: {
		ID:               242,
		Handle:           "nimiq",
		Symbol:           "NIM",
		Name:             "Nimiq",
		Decimals:         5,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	},
	IOTX: {
		ID:               304,
		Handle:           "iotex",
		Symbol:           "IOTX",
		Name:             "IoTeX",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
	},
	ZIL: {
		ID:               313,
		Handle:           "zilliqa",
		Symbol:           "ZIL",
		Name:             "Zilliqa",
		Decimals:         12,
		BlockTime:        30000,
		MinConfirmations: 1,
		SampleAddr:       "zil1anrjcsj2ntklaa3arq4w3s6gw4l4hqrycs9egy",
	},
	AION: {
		ID:               425,
		Handle:           "aion",
		Symbol:           "AION",
		Name:             "Aion",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0xa07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3",
	},
	AE: {
		ID:               457,
		Handle:           "aeternity",
		Symbol:           "AE",
		Name:             "Aeternity",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "ak_2p5878zbFhxnrm7meL7TmqwtvBaqcBddyp5eGzZbovZ5FeVfcw",
	},
	THETA: {
		ID:               500,
		Handle:           "theta",
		Symbol:           "THETA",
		Name:             "Theta",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
	},
	BNB: {
		ID:               714,
		Handle:           "binance",
		Symbol:           "BNB",
		Name:             "Binance",
		Decimals:         8,
		BlockTime:        4000,
		MinConfirmations: 0,
		SampleAddr:       "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	},
	VET: {
		ID:               818,
		Handle:           "vechain",
		Symbol:           "VET",
		Name:             "VeChain Token",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "0xB5e883349e68aB59307d1604555AC890fAC47128",
	},
	CLO: {
		ID:               820,
		Handle:           "callisto",
		Symbol:           "CLO",
		Name:             "Callisto",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0x39ec1c88a7a7c1a575e8c8f42eff7630d9278179",
	},
	TOMO: {
		ID:               889,
		Handle:           "tomochain",
		Symbol:           "TOMO",
		Name:             "TOMO",
		Decimals:         18,
		BlockTime:        4000,
		MinConfirmations: 0,
		SampleAddr:       "0x7daa83030e3086477b79b6e757ca8608899fe783",
	},
	TT: {
		ID:               1001,
		Handle:           "thundertoken",
		Symbol:           "TT",
		Name:             "ThunderCore",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0x0ad80a408eac4f17ba0a9de8a12d8736f60700c3",
	},
	ONT: {
		ID:               1024,
		Handle:           "ontology",
		Symbol:           "ONT",
		Name:             "Ontology",
		Decimals:         0,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	},
	XTZ: {
		ID:               1729,
		Handle:           "tezos",
		Symbol:           "XTZ",
		Name:             "Tezos",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
	},
	KIN: {
		ID:               2017,
		Handle:           "kin",
		Symbol:           "KIN",
		Name:             "Kin",
		Decimals:         5,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
	},
	NAS: {
		ID:               2718,
		Handle:           "nebulas",
		Symbol:           "NAS",
		Name:             "Nebulas",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a",
	},
	GO: {
		ID:               6060,
		Handle:           "gochain",
		Symbol:           "GO",
		Name:             "GoChain GO",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "0x76c2F81716A8D198a00502Ae9a59126418899FDe",
	},
	WAN: {
		ID:               5718350,
		Handle:           "wanchain",
		Symbol:           "WAN",
		Name:             "Wanchain",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "0x36cEdc3A9d969306AF4F7CA2b83ABBf74095914d",
	},
	WAVES: {
		ID:               5741564,
		Handle:           "waves",
		Symbol:           "WAVES",
		Name:             "WAVES",
		Decimals:         8,
		BlockTime:        30000,
		MinConfirmations: 1,
		SampleAddr:       "3P7wz6TXienpw3BHe8eHUEuZWb6WE58kgnQ",
	},
	SEM: {
		ID:               7562605,
		Handle:           "semux",
		Symbol:           "SEM",
		Name:             "Semux",
		Decimals:         9,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "0x8197987c401a3466ad678b2080b24838ebd95b41",
	},
	BTC: {
		ID:               0,
		Handle:           "bitcoin",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddr:       "bc1quvuarfksewfeuevuc6tn0kfyptgjvwsvrprk9d",
	},
	LTC: {
		ID:               2,
		Handle:           "litecoin",
		Symbol:           "LTC",
		Name:             "Litecoin",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
		SampleAddr:       "ltc1qhd8fxxp2dx3vsmpac43z6ev0kllm4n53t5sk0u",
	},
	DOGE: {
		ID:               3,
		Handle:           "doge",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "DJRFZNg8jkUtjcpo2zJd92FUAzwRjitw6f",
	},
	DASH: {
		ID:               5,
		Handle:           "dash",
		Symbol:           "DASH",
		Name:             "Dash",
		Decimals:         8,
		BlockTime:        180000,
		MinConfirmations: 0,
		SampleAddr:       "XqHiz8EXYbTAtBEYs4pWTHh7ipEDQcNQeT",
	},
	VIA: {
		ID:               14,
		Handle:           "viacoin",
		Symbol:           "VIA",
		Name:             "Viacoin",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
		SampleAddr:       "via1qnmsgjd6cvfprnszdgmyg9kewtjfgqflz67wwhc",
	},
	GRS: {
		ID:               17,
		Handle:           "groestlcoin",
		Symbol:           "GRS",
		Name:             "Groestlcoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "grs1qexwmshts5pdpeqglkl39zyl6693tmfwp0cue4j",
	},
	ZEC: {
		ID:               133,
		Handle:           "zcash",
		Symbol:           "ZEC",
		Name:             "Zcash",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
		SampleAddr:       "t1YYnByMzdGhQv3W3rnjHMrJs6HH4Y231gy",
	},
	XZC: {
		ID:               136,
		Handle:           "zcoin",
		Symbol:           "XZC",
		Name:             "Zcoin",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
		SampleAddr:       "aEd5XFChyXobvEics2ppAqgK3Bgusjxtik",
	},
	BCH: {
		ID:               145,
		Handle:           "bitcoincash",
		Symbol:           "BCH",
		Name:             "Bitcoin Cash",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddr:       "bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme",
	},
	RVN: {
		ID:               175,
		Handle:           "ravencoin",
		Symbol:           "RVN",
		Name:             "Raven",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "RHoCwPc2FCQqwToYnSiAb3SrCET4zEHsbS",
	},
	QTUM: {
		ID:               2301,
		Handle:           "qtum",
		Symbol:           "QTUM",
		Name:             "Qtum",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "QhceuaTdeCZtcxmVc6yyEDEJ7Riu5gWFoF",
	},
	ZEL: {
		ID:               19167,
		Handle:           "zelcash",
		Symbol:           "ZEL",
		Name:             "Zelcash",
		Decimals:         8,
		BlockTime:        120000,
		MinConfirmations: 0,
		SampleAddr:       "t1UKbRPzL4WN8Rs8aZ8RNiWoD2ftCMHKGUf",
	},
	DCR: {
		ID:               42,
		Handle:           "decred",
		Symbol:           "DCR",
		Name:             "Decred",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
		SampleAddr:       "DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY",
	},
}
func Ethereum() Coin {
	return Coins[ETH]
}
func Classic() Coin {
	return Coins[ETC]
}
func Icon() Coin {
	return Coins[ICX]
}
func Cosmos() Coin {
	return Coins[ATOM]
}
func Ripple() Coin {
	return Coins[XRP]
}
func Stellar() Coin {
	return Coins[XLM]
}
func Poa() Coin {
	return Coins[POA]
}
func Tron() Coin {
	return Coins[TRX]
}
func Fio() Coin {
	return Coins[FIO]
}
func Nimiq() Coin {
	return Coins[NIM]
}
func Iotex() Coin {
	return Coins[IOTX]
}
func Zilliqa() Coin {
	return Coins[ZIL]
}
func Aion() Coin {
	return Coins[AION]
}
func Aeternity() Coin {
	return Coins[AE]
}
func Theta() Coin {
	return Coins[THETA]
}
func Binance() Coin {
	return Coins[BNB]
}
func Vechain() Coin {
	return Coins[VET]
}
func Callisto() Coin {
	return Coins[CLO]
}
func Tomochain() Coin {
	return Coins[TOMO]
}
func Thundertoken() Coin {
	return Coins[TT]
}
func Ontology() Coin {
	return Coins[ONT]
}
func Tezos() Coin {
	return Coins[XTZ]
}
func Kin() Coin {
	return Coins[KIN]
}
func Nebulas() Coin {
	return Coins[NAS]
}
func Gochain() Coin {
	return Coins[GO]
}
func Wanchain() Coin {
	return Coins[WAN]
}
func Waves() Coin {
	return Coins[WAVES]
}
func Semux() Coin {
	return Coins[SEM]
}
func Bitcoin() Coin {
	return Coins[BTC]
}
func Litecoin() Coin {
	return Coins[LTC]
}
func Doge() Coin {
	return Coins[DOGE]
}
func Dash() Coin {
	return Coins[DASH]
}
func Viacoin() Coin {
	return Coins[VIA]
}
func Groestlcoin() Coin {
	return Coins[GRS]
}
func Zcash() Coin {
	return Coins[ZEC]
}
func Zcoin() Coin {
	return Coins[XZC]
}
func Bitcoincash() Coin {
	return Coins[BCH]
}
func Ravencoin() Coin {
	return Coins[RVN]
}
func Qtum() Coin {
	return Coins[QTUM]
}
func Zelcash() Coin {
	return Coins[ZEL]
}
func Decred() Coin {
	return Coins[DCR]
}

