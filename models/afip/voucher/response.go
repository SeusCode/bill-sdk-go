package voucher

import "github.com/seuscode/bill-sdk-go/models/afip/document"

type (
	AfipInformation struct {
		VoucherDate  string                `json:"CbteFch"`
		VoucherType  VoucherType           `json:"CbteTipo"`
		Concept      VoucherConcept        `json:"Concepto"`
		Document     int64                 `json:"DocNro"`
		DocumentType document.DocumentType `json:"DocTipo"`
		ImpIVA       float64               `json:"ImpIVA"`
		ImpNeto      float64               `json:"ImpNeto"`
		ImpOpEx      float64               `json:"ImpOpEx"`
		ImpTotal     float64               `json:"ImpTotal"`
		ImpTotConc   float64               `json:"ImpTotConc"`
		ImpTrib      float64               `json:"ImpTrib"`
		MonCotiz     float64               `json:"MonCotiz"`
		MonId        string                `json:"MonId"`
		PointOfSale  int                   `json:"PtoVta"`
		Iva          []VoucherShare        `json:"Iva"`
	}

	CreateVoucherResponse struct {
		CaeNumber         string          `json:"CAE"`
		VoucherNumber     int             `json:"voucher_number"`
		CaeExpirationDate string          `json:"CAEFchVto"`
		AfipInfo          AfipInformation `json:"info_to_afip"`
		QrEndpoint        string          `json:"qr_endpoint"`
	}

	GetVoucherTypesResponse struct {
		Vouchers []Type `json:"voucher_types"`
	}
)
