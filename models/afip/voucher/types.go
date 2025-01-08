package voucher

import (
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
)

type (
	VoucherType    int32
	VoucherConcept int32

	Type struct {
		Id        VoucherType `json:"Id"`
		Desc      string      `json:"Desc"`
		StartDate string      `json:"FchDesde"`
		EndDate   string      `json:"FchHasta"`
	}

	Voucher struct {
		CbteTipo           VoucherType           `json:"CbteTipo"`                     // Tipo de Comprobante
		Concepto           VoucherConcept        `json:"Concepto"`                     // Concepto del Comprobante
		DocTipo            document.DocumentType `json:"DocTipo"`                      // Tipo de Documento del Receptor
		DocNro             int64                 `json:"DocNro"`                       // Número de Documento del Receptor
		CbteFch            *string               `json:"CbteFch,omitempty"`            // Fecha del Comprobante (opcional: puede ser actual si no se envía)
		MonId              *string               `json:"MonId,omitempty"`              // Moneda (por defecto "PES")
		MonCotiz           *float64              `json:"MonCotiz,omitempty"`           // Cotización de la Moneda (por defecto 1)
		FchServDesde       *string               `json:"FchServDesde,omitempty"`       // Fecha de Inicio del Servicio (requerido para Concepto 2 o 3)
		FchServHasta       *string               `json:"FchServHasta,omitempty"`       // Fecha de Fin del Servicio (requerido para Concepto 2 o 3)
		FchVtoPago         *string               `json:"FchVtoPago,omitempty"`         // Fecha de Vencimiento del Pago (requerido para Concepto 2 o 3)
		Items              []VoucherItems        `json:"Items"`                        // Detalle de los ítems del comprobante
		CbtesAsoc          []AsociatedVouchers   `json:"CbtesAsoc,omitempty"`          // Comprobantes Asociados (opcional)
		Tributos           []VoucherTributes     `json:"Tributos,omitempty"`           // Tributos (opcional)
		Opcionales         []VoucherOptionals    `json:"Opcionales,omitempty"`         // Opcionales (opcional)
		Compradores        []VoucherBuyers       `json:"Compradores,omitempty"`        // Compradores adicionales (opcional)
		CompradorIvaExento *bool                 `json:"CompradorIvaExento,omitempty"` // Indica si el comprador es IVA exento (opcional)
		UniqueItems        *bool                 `json:"UniqueItems,omitempty"`        // Indica si los ítems son únicos (opcional)
	}

	VoucherItems struct {
		Id             string            `json:"Id"`
		Qty            float64           `json:"Qty"`
		Iva            aliquot.AliquotId `json:"Iva"`
		Price          float64           `json:"Price"`
		Desc           string            `json:"Desc"`
		Discount       float64           `json:"DiscountPercent"`
		Subtotal       float64           `json:"Subtotal"`
		IvaExento      bool              `json:"IvaExento"`
		NoGravadodeIVA bool              `json:"NoGravadodeIVA"`
	}

	VoucherTributes struct {
		Id      int     `json:"Id"`
		Desc    string  `json:"Desc"`
		BaseImp float64 `json:"BaseImp"`
		Alic    float64 `json:"Alic"`
		Total   float64 `json:"Importe"`
	}

	VoucherOptionals struct {
		Id    int `json:"Id"`
		Value int `json:"Valor"`
	}

	VoucherShare struct {
		Id      int     `json:"Id"`
		BaseImp float64 `json:"BaseImp"`
		Total   float64 `json:"Importe"`
	}

	VoucherBuyers struct {
		DocType    document.DocumentType `json:"DocTipo"`
		DocNro     int                   `json:"DocNro"`
		Percentage int                   `json:"Porcentaje"`
	}

	AsociatedVouchers struct {
		Type  VoucherType `json:"Tipo"`
		Pos   int         `json:"PtoVta"`
		Nbr   int         `json:"Nro"`
		TaxId int         `json:"Cuit"`
	}
)

const (
	FacturaA VoucherType = 1
	FacturaB VoucherType = 6
	FacturaC VoucherType = 11

	Ambos     VoucherConcept = 3
	Productos VoucherConcept = 1
	Servicios VoucherConcept = 2
)
