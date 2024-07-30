package voucher

import (
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
)

type (
	VoucherType        uint
	VoucherConcept     uint
	VoucherPDFTemplate string

	Type struct {
		Id        VoucherType `json:"Id"`
		Desc      string      `json:"Desc"`
		StartDate string      `json:"FchDesde"`
		EndDate   string      `json:"FchHasta"`
	}

	Voucher struct {
		CbteTipo VoucherType    `json:"CbteTipo"`
		Concepto VoucherConcept `json:"Concepto"`

		DocTipo document.DocumentType `json:"DocTipo"`
		DocNro  int64                 `json:"DocNro"`

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

		MonId    *string  `json:"MonId,omitempty"`
		MonCotiz *float64 `json:"MonCotiz,omitempty"`

		Phone *string `json:"tel,omitempty"`
		Email *string `json:"email,omitempty"`
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

	VoucherClient struct {
		Phone              string                         `json:"tel,omitempty"`
		Email              string                         `json:"email,omitempty"`
		Address            string                         `json:"domicilio"`
		FiscalType         fiscal.FiscalType              `json:"condicionIva"`
		SellCondition      payment_method.SellConditionId `json:"condicionDeVenta"`
		NameOrBusinessName string                         `json:"razonSocial"`
	}

	VoucherPDF struct {
		Logo      string             `json:"logo"`
		Watermark string             `json:"watermark"`
		Template  VoucherPDFTemplate `json:"template"`
		PtoVta    int                `json:"PtoVta"`

		CbteNro  int            `json:"CbteNum"`
		CbteFch  int64          `json:"CbteFch"`
		CbteTipo VoucherType    `json:"CbteTipo"`
		Concepto VoucherConcept `json:"Concepto"`

		CopiaDesde int `json:"fromCopy"`
		CopiaHasta int `json:"toCopy"`

		DocTipo document.DocumentType `json:"DocTipo"`
		DocNro  int64                 `json:"DocNro"`

		MonId    string  `json:"MonId"`
		MonCotiz float64 `json:"MonCotiz"`

		CAE       string `json:"CAE"`
		CAEFchVto string `json:"CAEFchVto"`

		Cliente VoucherClient `json:"Cliente"`

		ImpIVA     float64 `json:"ImpIVA"`
		ImpNeto    float64 `json:"ImpNeto"`
		ImpOpEx    float64 `json:"ImpOpEx"`
		ImpTrib    float64 `json:"ImpTrib"`
		ImpTotal   float64 `json:"ImpTotal"`
		ImpTotConc float64 `json:"ImpTotConc"`

		Iva         []VoucherShare      `json:"Iva,omitempty"`
		CbtesAsoc   []AsociatedVouchers `json:"CbtesAsoc,omitempty"`
		Tributos    []VoucherTributes   `json:"Tributos,omitempty"`
		Opcionales  []VoucherOptionals  `json:"Opcionales,omitempty"`
		Compradores []VoucherBuyers     `json:"Compradores,omitempty"`

		Items []VoucherItems `json:"Items"`

		FchServDesde string `json:"FchServDesde"`
		FchServHasta string `json:"FchServHasta"`
		FchVtoPago   string `json:"FchVtoPago"`
	}
)

const (
	FacturaA VoucherType = 1
	FacturaB VoucherType = 6
	FacturaC VoucherType = 11

	Ambos     VoucherConcept = 3
	Productos VoucherConcept = 1
	Servicios VoucherConcept = 2

	Clasico VoucherPDFTemplate = "classic"
)
