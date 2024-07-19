package voucher

import "github.com/seuscode/bill-sdk-go/models/afip/document"

type (
	CreateVoucherRequest struct {
		PtoVta int `json:"PtoVta"`

		CbteTipo VoucherType    `json:"CbteTipo"`
		Concepto VoucherConcept `json:"Concepto"`

		DocTipo document.DocumentType `json:"DocTipo"`
		DocNro  int                   `json:"DocNro"`

		CbteFch *string `json:"CbteFch,omitempty"`

		FchServDesde *int `json:"FchServDesde,omitempty"`
		FchServHasta *int `json:"FchServHasta,omitempty"`
		FchVtoPago   *int `json:"FchVtoPago,omitempty"`

		Items []VoucherItems `json:"Items"`

		CbtesAsoc   []AsociatedVouchers `json:"CbtesAsoc,omitempty"`
		Iva         []VoucherShare      `json:"Iva,omitempty"`
		Tributos    []VoucherTributes   `json:"Tributos,omitempty"`
		Opcionales  []VoucherOptionals  `json:"Opcionales,omitempty"`
		Compradores []VoucherBuyers     `json:"Compradores,omitempty"`

		CompradorIvaExento *bool `json:"CompradorIvaExento,omitempty"`
		PagoContado        *bool `json:"pagoContado,omitempty"`
		GeneratePDF        *bool `json:"doPDF,omitempty"`

		MonId    *string  `json:"MonId,omitempty"`
		MonCotiz *float64 `json:"MonCotiz,omitempty"`
	}
)
