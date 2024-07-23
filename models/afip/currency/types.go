package currency

type (
	Currency struct {
		Id        string `json:"Id"`
		Desc      string `json:"Desc"`
		StartDate string `json:"FchDesde"`
		EndDate   string `json:"FchHasta"`
	}

	Cotization struct {
		CotizationDate  string  `json:"FchCotiz"`
		CotizationPrice float64 `json:"MonCotiz"`
		CurrencyId      string  `json:"MonId"`
		Description     string  `json:"description"`
		Full            string  `json:"full"`
		Short           string  `json:"short"`
		Timestamp       int64   `json:"timestamp"`
	}
)
