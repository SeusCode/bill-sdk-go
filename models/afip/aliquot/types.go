package aliquot

type (
	AliquotId string

	Aliquot struct {
		Id         AliquotId `json:"Id"`
		EndDate    string    `json:"FchHasta"`
		StartDate  string    `json:"FchDesde"`
		Percentage string    `json:"Desc"`
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
