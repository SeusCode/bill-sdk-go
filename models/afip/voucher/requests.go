package voucher

import (
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
)

type (
	CreateVoucherRequest struct {
		CbteTipo           VoucherType           `json:"CbteTipo"`
		Concepto           VoucherConcept        `json:"Concepto"`
		DocTipo            document.DocumentType `json:"DocTipo"`
		DocNro             int64                 `json:"DocNro"`
		CbteFch            *string               `json:"CbteFch,omitempty"`
		MonId              *string               `json:"MonId,omitempty"`
		MonCotiz           *float64              `json:"MonCotiz,omitempty"`
		FchServDesde       *string               `json:"FchServDesde,omitempty"`
		FchServHasta       *string               `json:"FchServHasta,omitempty"`
		FchVtoPago         *string               `json:"FchVtoPago,omitempty"`
		Items              []VoucherItems        `json:"Items"`
		CbtesAsoc          []AsociatedVouchers   `json:"CbtesAsoc,omitempty"`
		Tributos           []VoucherTributes     `json:"Tributos,omitempty"`
		Opcionales         []VoucherOptionals    `json:"Opcionales,omitempty"`
		Compradores        []VoucherBuyers       `json:"Compradores,omitempty"`
		CompradorIvaExento *bool                 `json:"CompradorIvaExento,omitempty"`
		UniqueItems        *bool                 `json:"UniqueItems,omitempty"`
	}

	CreateVoucherPDFRequest struct {
		CAE                      string                         `json:"CAE"`
		CAEFchVto                string                         `json:"CAEFchVto"`
		QREndpoint               string                         `json:"qr_endpoint"`
		AfipInformation          AFIPInformation                `json:"info_to_afip"`
		VoucherNbr               int32                          `json:"voucher_number"`
		MetodoPago               payment_method.SellConditionId `json:"metodo_pago"`
		ClientFiscalType         uint                           `json:"client_fiscal_type"`
		ClientPhone              string                         `json:"client_phone"`
		ClientEmail              string                         `json:"client_email"`
		ClientAddress            string                         `json:"client_address"`
		ClientNameOrBusinessName string                         `json:"client_name"`
	}
)
