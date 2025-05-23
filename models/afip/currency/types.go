package currency

type (
	CurrencyDetails struct {
		Id       string `xml:"Id,omitempty" json:"Id,omitempty"`
		Desc     string `xml:"Desc,omitempty" json:"Desc,omitempty"`
		FchDesde string `xml:"FchDesde,omitempty" json:"FchDesde,omitempty"`
		FchHasta string `xml:"FchHasta,omitempty" json:"FchHasta,omitempty"`
	}

	CurrencyExchangeRate struct {
		MonId    string  `xml:"MonId,omitempty" json:"MonId,omitempty"`
		MonCotiz float64 `xml:"MonCotiz,omitempty" json:"MonCotiz,omitempty"`
		FchCotiz string  `xml:"FchCotiz,omitempty" json:"FchCotiz,omitempty"`
	}
)
