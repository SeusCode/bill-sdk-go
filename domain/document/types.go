package document

type (
	DocumentType uint

	Document struct {
		Id        DocumentType `json:"Id"`
		Desc      string       `json:"Desc"`
		StartDate string       `json:"FchDesde"`
		EndDate   string       `json:"FchHasta"`
	}
)

const (
	CF   DocumentType = 99 // Consumidor final
	DNI  DocumentType = 96
	CUIL DocumentType = 86
	CUIT DocumentType = 80
)
