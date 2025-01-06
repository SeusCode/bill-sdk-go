package currency

type (
	GetCurrencyTypesResponse struct {
		Currencies []Currency `json:"currency_types"`
	}

	GetCurrencyCotizationResponse struct {
		CotizationData Cotization `json:"currency_quotation"`
	}
)
