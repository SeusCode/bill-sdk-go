package responses

import "github.com/seuscode/bill-sdk-go/domain/tax"

type (
	GetTaxTypesResponse struct {
		Taxes []tax.Tax `json:"tax_types"`
	}
)
