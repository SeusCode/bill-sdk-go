package requests

import (
	"github.com/seuscode/afip-sdk-go/domain/document"
	"github.com/seuscode/afip-sdk-go/domain/voucher"
)

type (
	CreateVoucher struct {
		PtoVta int `json:"PtoVta"`

		CbteTipo voucher.VoucherType    `json:"CbteTipo"`
		Concepto voucher.VoucherConcept `json:"Concepto"`

		DocTipo document.DocumentType `json:"DocTipo"`
		DocNro  int                   `json:"DocNro"`

		CbteFch *string `json:"CbteFch,omitempty"`

		FchServDesde *int `json:"FchServDesde,omitempty"`
		FchServHasta *int `json:"FchServHasta,omitempty"`
		FchVtoPago   *int `json:"FchVtoPago,omitempty"`

		Items []voucher.VoucherItems `json:"Items"`

		CbtesAsoc   []voucher.AsociatedVouchers `json:"CbtesAsoc,omitempty"`
		Iva         []voucher.VoucherShare      `json:"Iva,omitempty"`
		Tributos    []voucher.VoucherTributes   `json:"Tributos,omitempty"`
		Opcionales  []voucher.VoucherOptionals  `json:"Opcionales,omitempty"`
		Compradores []voucher.VoucherBuyers     `json:"Compradores,omitempty"`

		CompradorIvaExento *bool `json:"CompradorIvaExento,omitempty"`
		PagoContado        *bool `json:"pagoContado,omitempty"`
		GeneratePDF        *bool `json:"doPDF,omitempty"`

		MonId    *string  `json:"MonId,omitempty"`
		MonCotiz *float64 `json:"MonCotiz,omitempty"`
	}
)
