package responses

import "github.com/seuscode/bill-sdk-go/domain/aliquot"

type (
	GetAliquotTypesResponse struct {
		Aliquots []aliquot.Aliquot `json:"aliquot_types"`
	}
)
