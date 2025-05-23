package document

type (
	DocumentType int32

	DocumentDetails struct {
		Id       DocumentType `xml:"Id,omitempty" json:"Id,omitempty"`
		Desc     string       `xml:"Desc,omitempty" json:"Desc,omitempty"`
		FchDesde string       `xml:"FchDesde,omitempty" json:"FchDesde,omitempty"`
		FchHasta string       `xml:"FchHasta,omitempty" json:"FchHasta,omitempty"`
	}
)

const (
	CF   DocumentType = 99 // Consumidor final
	DNI  DocumentType = 96
	CUIL DocumentType = 86
	CUIT DocumentType = 80
)
