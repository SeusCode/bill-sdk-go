package invoice

import "github.com/seuscode/bill-sdk-go/models/afip/document"

type (
	IssueInvoiceRequest struct {
		CbteTipo     InvoiceType                `json:"CbteTipo" uker:"required"` // Tipo de Comprobante
		Concepto     InvoiceConcept             `json:"Concepto" uker:"required"` // Concepto del Comprobante
		DocTipo      document.DocumentType      `json:"DocTipo" uker:"required"`  // Tipo de Documento del Receptor
		DocNro       int64                      `json:"DocNro" uker:"required"`   // Número de Documento del Receptor
		CondFisc     uint                       `json:"CondFisc" uker:"required"` // Condicion fiscal del comprador
		CbteFch      *string                    `json:"CbteFch,omitempty"`        // Fecha del Comprobante (opcional: puede ser actual si no se envía)
		MonId        *string                    `json:"MonId,omitempty"`          // Moneda (por defecto "PES")
		MonCotiz     *float64                   `json:"MonCotiz,omitempty"`       // Cotización de la Moneda (por defecto 1)
		FchServDesde *string                    `json:"FchServDesde,omitempty"`   // Fecha de Inicio del Servicio (requerido para Concepto 2 o 3)
		FchServHasta *string                    `json:"FchServHasta,omitempty"`   // Fecha de Fin del Servicio (requerido para Concepto 2 o 3)
		FchVtoPago   *string                    `json:"FchVtoPago,omitempty"`     // Fecha de Vencimiento del Pago (requerido para Concepto 2 o 3)
		CanMisMonExt *string                    `json:"CanMisMonExt,omitempty"`   //  Marca que identifica si el comprobante se cancela en misma moneda del comprobante. (opcional)
		Items        []InvoiceItem              `json:"Items" uker:"required"`    // Detalle de los ítems del comprobante
		CbtesAsoc    []InvoiceAssociatedVoucher `json:"CbtesAsoc,omitempty"`      // Comprobantes Asociados (opcional)
		Tributos     []InvoiceTribute           `json:"Tributos,omitempty"`       // Tributos (opcional)
		Opcionales   []InvoiceOptional          `json:"Opcionales,omitempty"`     // Opcionales (opcional)
		Compradores  []InvoiceBuyer             `json:"Compradores,omitempty"`    // Compradores adicionales (opcional)
		ItemsUnicos  *bool                      `json:"ItemsUnicos,omitempty"`    // Indica si los ítems son únicos (opcional)
	}

	GenerateInvoicePDFRequest struct {
		InvoiceData     IssueInvoiceResponse `json:"data" uker:"required"`
		InvoiceMetaData InvoiceMetaData      `json:"metadata" uker:"required"`
		InvoiceProducts []InvoiceItem        `json:"products" uker:"required"`
	}
)
