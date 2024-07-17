package responses

import "github.com/seuscode/bill-sdk-go/domain/document"

type (
	GetDocumentTypesResponse struct {
		Documents []document.Document `json:"document_types"`
	}
)
