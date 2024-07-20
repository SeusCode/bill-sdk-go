package citizen

type (
	GetPersonInformationRequest struct {
		TaxId          int64 `json:"cuit"`
		CitizenId      int64 `json:"documento,omitempty"`
		RegistryNumber int   `json:"padron"`
	}
)
