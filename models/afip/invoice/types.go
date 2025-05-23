package invoice

import (
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
)

type (
	InvoiceType       int
	InvoiceConcept    int
	InvoiceFiscalType uint

	InvoiceMetaData struct {
		PaymentMethod            payment_method.SellConditionId `json:"metodo_pago,omitempty" uker:"required"`
		IndirectNacionalTaxTotal float64                        `json:"total_oini,omitempty"` // Total de impuestos indirectos nacionales
		ClientFiscalType         InvoiceFiscalType              `json:"client_fiscal_type,omitempty" uker:"required"`
		ClientPhone              string                         `json:"client_phone,omitempty"`
		ClientEmail              string                         `json:"client_email,omitempty"`
		ClientAddress            string                         `json:"client_address,omitempty"`
		ClientNameOrBusinessName string                         `json:"client_name,omitempty"`
	}

	InvoiceItem struct {
		Id             string            `json:"Id" uker:"required"`
		Desc           string            `json:"Desc" uker:"required"`
		Precio         float64           `json:"Precio" uker:"required"`
		Cantidad       float64           `json:"Cantidad" uker:"required"`
		Iva            aliquot.AliquotId `json:"Iva" uker:"required"`
		IvaExento      *bool             `json:"IvaExento,omitempty"`
		NoGravadodeIVA *bool             `json:"NoGravadodeIVA,omitempty"`
		Subtotal       float64           `json:"Subtotal" uker:"required"`
		MontoIva       float64           `json:"MontoIva,omitempty"`
		TotalCIva      float64           `json:"TotalCIva,omitempty"`
		PriceCIva      float64           `json:"PriceCIva,omitempty"`
	}

	InvoiceAssociatedVoucher struct {
		Nro    int   `json:"Nro" uker:"required"`
		Tipo   int   `json:"Tipo" uker:"required"`
		Cuit   int64 `json:"Cuit" uker:"required"`
		PtoVta int   `json:"PtoVta" uker:"required"`
	}

	InvoiceTribute struct {
		Id      int     `json:"Id" uker:"required"`
		Desc    string  `json:"Desc" uker:"required"`
		BaseImp float64 `json:"BaseImp" uker:"required"`
		Alic    float64 `json:"Alic" uker:"required"`
		Importe float64 `json:"Importe" uker:"required"`
	}

	InvoiceOptional struct {
		Id    int     `json:"Id" uker:"required"`
		Valor float64 `json:"Valor" uker:"required"`
	}

	InvoiceBuyer struct {
		DocTipo    int     `json:"DocTipo" uker:"required"`
		DocNro     int64   `json:"DocNro" uker:"required"`
		Porcentaje float64 `json:"Porcentaje" uker:"required"`
	}

	InvoiceTax struct {
		Id      int     `json:"Id" uker:"required"`
		BaseImp float64 `json:"BaseImp" uker:"required"`
		Importe float64 `json:"Importe" uker:"required"`
	}

	InvoicePeriod struct {
		FchDesde string `json:"FchDesde" uker:"required"`
		FchHasta string `json:"FchHasta" uker:"required"`
	}

	InvoiceActivity struct {
		Id int64 `json:"Id" uker:"required"`
	}

	InvoiceObservation struct {
		Code int32  `xml:"Code,omitempty" json:"Code,omitempty"`
		Msg  string `xml:"Msg,omitempty" json:"Msg,omitempty"`
	}
)

const (
	// Fiscal Types
	MONOTRIBUTO                  InvoiceFiscalType = 1
	NO_CATEGORIZADO              InvoiceFiscalType = 2
	CONSUMIDOR_FINAL             InvoiceFiscalType = 3
	PROVEEDOR_DEL_EXTERIOR       InvoiceFiscalType = 4
	IVA_EXENTO                   InvoiceFiscalType = 5
	IVA_RESPONSABLE_INSCRIPTO    InvoiceFiscalType = 6
	IVA_RESPONSABLE_NO_INSCRIPTO InvoiceFiscalType = 7

	FACTURA_A InvoiceType = 1
	FACTURA_B InvoiceType = 6
	FACTURA_C InvoiceType = 11

	AMBOS     InvoiceConcept = 3
	PRODUCTOS InvoiceConcept = 1
	SERVICIOS InvoiceConcept = 2
)
