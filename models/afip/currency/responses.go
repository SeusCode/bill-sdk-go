package currency

type (
	GetCurrenciesResponse struct {
		Currencies []CurrencyDetails `json:"currencies"`
	}

	GetCurrencyExchangeRateResponse struct {
		ExchangeRate CurrencyExchangeRate `json:"exchange_rate"`
	}
)
