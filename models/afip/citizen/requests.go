package citizen

type (
	GetPersonInformationRequest struct {
		TaxId          string `json:"cuit"`
		CitizenId      string `json:"documento,omitempty"`
		RegistryNumber int    `json:"padron"`
	}
)
