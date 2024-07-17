package pos

type (
	PointOfSale struct {
		Id              int    `json:"Nro"`
		Type            string `json:"EmisionTipo"`
		Banned          string `json:"Bloqueado"`
		TerminationDate string `json:"FchBaja"`
	}
)
