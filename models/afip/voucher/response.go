package voucher

import "github.com/seuscode/bill-sdk-go/models/afip/document"

type (
	AFIPInformation struct {
		VoucherDate  string              `json:"CbteFch"`
		VoucherType  int32               `json:"CbteTipo"`
		Concept      int32               `json:"Concepto"`
		Document     int64               `json:"DocNro"`
		DocumentType int32               `json:"DocTipo"`
		ImpIVA       float64             `json:"ImpIVA"`
		ImpNeto      float64             `json:"ImpNeto"`
		ImpOpEx      float64             `json:"ImpOpEx"`
		ImpTotal     float64             `json:"ImpTotal"`
		ImpTotConc   float64             `json:"ImpTotConc"`
		ImpTrib      float64             `json:"ImpTrib"`
		MonCotiz     float64             `json:"MonCotiz"`
		MonId        string              `json:"MonId"`
		PointOfSale  int32               `json:"PtoVta"`
		FchServDesde string              `json:"FchServDesde,omitempty"`
		FchServHasta string              `json:"FchServHasta,omitempty"`
		FchVtoPago   string              `json:"FchVtoPago,omitempty"`
		Iva          []VoucherShare      `json:"Iva,omitempty"`
		Items        []VoucherItems      `json:"Items"`
		CbtesAsoc    []AsociatedVouchers `json:"CbtesAsoc,omitempty"`
		Tributos     []VoucherTributes   `json:"Tributos,omitempty"`
		Opcionales   []VoucherOptionals  `json:"Opcionales,omitempty"`
		Compradores  []VoucherBuyers     `json:"Compradores,omitempty"`
	}

	AFIPVoucherIva struct {
		AlicIva VoucherShare `json:"AlicIva"`
	}

	AFIPVoucherData struct {
		CbteDesde       int                   `json:"CbteDesde"`
		CbteFch         string                `json:"CbteFch"`
		CbteHasta       int                   `json:"CbteHasta"`
		CbteTipo        VoucherType           `json:"CbteTipo"`
		CodAutorizacion string                `json:"CodAutorizacion"`
		Concepto        VoucherConcept        `json:"Concepto"`
		DocNro          int64                 `json:"DocNro"`
		DocTipo         document.DocumentType `json:"DocTipo"`
		EmisionTipo     string                `json:"EmisionTipo"`
		FchProceso      string                `json:"FchProceso"`
		FchServDesde    string                `json:"FchServDesde"`
		FchServHasta    string                `json:"FchServHasta"`
		FchVto          string                `json:"FchVto"`
		FchVtoPago      string                `json:"FchVtoPago"`
		ImpIVA          float64               `json:"ImpIVA"`
		ImpNeto         float64               `json:"ImpNeto"`
		ImpOpEx         float64               `json:"ImpOpEx"`
		ImpTotConc      float64               `json:"ImpTotConc"`
		ImpTotal        float64               `json:"ImpTotal"`
		ImpTrib         float64               `json:"ImpTrib"`
		Iva             AFIPVoucherIva        `json:"Iva"`
		MonCotiz        float64               `json:"MonCotiz"`
		MonId           string                `json:"MonId"`
		PtoVta          int                   `json:"PtoVta"`
		Resultado       string                `json:"Resultado"`
	}

	CreateVoucherResponse struct {
		CaeNumber         string          `json:"CAE"`
		VoucherNumber     int             `json:"voucher_number"`
		CaeExpirationDate string          `json:"CAEFchVto"`
		AfipInfo          AFIPInformation `json:"info_to_afip"`
		QrEndpoint        string          `json:"qr_endpoint"`
	}

	GetVoucherTypesResponse struct {
		Vouchers []Type `json:"voucher_types"`
	}

	GetVoucherInfoResponse struct {
		Data        AFIPVoucherData `json:"data"`
		VoucherType string          `json:"CbteTipo"`
	}
)
