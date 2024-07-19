package tax

type (
	Tax struct {
		Id        int    `json:"Id"`
		Desc      string `json:"Desc"`
		StartDate string `json:"FchDesde"`
		EndDate   string `json:"FchHasta"`
	}
)
