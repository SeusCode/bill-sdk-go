package voucher

import (
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
)

type (
	VoucherType    uint
	VoucherConcept uint

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
		GeneratePDF        *bool `json:"doPDF,omitempty"`

		MonId    *string  `json:"MonId,omitempty"`
		MonCotiz *float64 `json:"MonCotiz,omitempty"`

		Phone   *string `json:"tel,omitempty"`
		Email   *string `json:"email,omitempty"`
		Website *string `json:"MonCowebtiz,omitempty"`
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
