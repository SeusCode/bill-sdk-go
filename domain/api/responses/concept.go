package responses

import "github.com/seuscode/bill-sdk-go/domain/concept"

type (
	GetConceptTypesResponse struct {
		Concepts []concept.Concept `json:"concept_types"`
	}
)
