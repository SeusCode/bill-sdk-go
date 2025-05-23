package aliquot

type (
	AliquotId string

	InvoiceAliquot struct {
		Id         AliquotId `json:"Id"`
		EndDate    string    `json:"FchHasta"`
		StartDate  string    `json:"FchDesde"`
		Percentage string    `json:"Desc"`
	}

	AliquotDetails struct {
		Id       string `xml:"Id,omitempty" json:"Id,omitempty"`
		Desc     string `xml:"Desc,omitempty" json:"Desc,omitempty"`
		FchDesde string `xml:"FchDesde,omitempty" json:"FchDesde,omitempty"`
		FchHasta string `xml:"FchHasta,omitempty" json:"FchHasta,omitempty"`
	}
)

const (
	ZeroPercent        AliquotId = "3"
	TwoDotFivePercent  AliquotId = "9"
	FivePercent        AliquotId = "8"
	TenDotFivePercent  AliquotId = "4"
	TwentyOnePercent   AliquotId = "5"
	TwentySevenPercent AliquotId = "6"
)
