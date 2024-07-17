package currency

type (
	Currency struct {
		Id        string `json:"Id"`
		Desc      string `json:"Desc"`
		StartDate string `json:"FchDesde"`
		EndDate   string `json:"FchHasta"`
	}
)
