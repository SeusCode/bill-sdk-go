package citizen

type (
	GetCitizenRecordResponse struct {
		CitizenP4  *CitizenP4  `json:"p4_citizen,omitempty"`
		CitizenP5  *CitizenP5  `json:"p5_citizen,omitempty"`
		CitizenP10 *CitizenP10 `json:"p10_citizen,omitempty"`
		CitizenP13 *CitizenP13 `json:"p13_citizen,omitempty"`
	}
)
