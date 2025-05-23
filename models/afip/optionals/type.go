package optionals

type (
	OptionalDetails struct {
		Id       string `xml:"Id,omitempty" json:"Id,omitempty"`
		Desc     string `xml:"Desc,omitempty" json:"Desc,omitempty"`
		FchDesde string `xml:"FchDesde,omitempty" json:"FchDesde,omitempty"`
		FchHasta string `xml:"FchHasta,omitempty" json:"FchHasta,omitempty"`
	}
)
