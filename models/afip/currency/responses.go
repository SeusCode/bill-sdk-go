package currency

type (
	GetCurrencyTypesResponse struct {
		Currencies []Currency `json:"currencies_types"`
	}

	GetCurrencyCotizationResponse struct {
		CotizationData Cotization `json:"monCotiz"`
	}
)
