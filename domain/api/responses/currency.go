package responses

import (
	"github.com/seuscode/bill-sdk-go/domain/currency"
)

type (
	GetCurrencyTypesResponse struct {
		Currencies []currency.Currency `json:"currencies_types"`
	}
)
