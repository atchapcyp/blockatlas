package assets

import (
	"github.com/trustwallet/blockatlas"
	"time"

	"github.com/trustwallet/blockatlas/coin"
	"net/http"
)

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func GetValidators(coin coin.Coin) ([]AssetValidator, error) {
	var results []AssetValidator
	request := blockatlas.Request{
		HttpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		ErrorHandler: func(res *http.Response, uri string) error {
			return nil
		},
	}
	err := request.Get(&results, AssetsURL+coin.Handle, "/validators/list.json", nil)
	return results, err
}

func NormalizeValidators(validators []blockatlas.Validator, assets []AssetValidator) []blockatlas.StakeValidator {
	results := make([]blockatlas.StakeValidator, 0)

	for _, v := range validators {
		for _, v2 := range assets {
			if v.ID == v2.ID {
				results = append(results, NormalizeValidator(v, v2))
			}
		}
	}

	return results
}

func NormalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator) blockatlas.StakeValidator {
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status,
		Info: blockatlas.StakeValidatorInfo{
			Name:        validator.Name,
			Description: validator.Description,
			Image:       GetImage(plainValidator.Coin, plainValidator.ID),
			Website:     validator.Website,
		},
		Reward: plainValidator.Reward,
	}
}

func GetImage(c coin.Coin, ID string) string {
	return AssetsURL + c.Handle + "/validators/assets/" + ID + "/logo.png"
}
