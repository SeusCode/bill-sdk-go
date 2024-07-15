package responses

import "github.com/seuscode/afip-sdk-go/domain/aliquot"

type (
	GetAliquotTypesResponse struct {
		Aliquots []aliquot.Aliquot `json:"aliquot_types"`
	}
)
