package invoice

type (
	IssueInvoiceResponse struct {
		CAE       string `json:"CAE"`
		CAEFchVto string `json:"CAEFchVto"`
		CbteFch   string `json:"CbteFch"`

		CbteTipo int32 `json:"CbteTipo"`
		PtoVta   int32 `json:"PtoVta"`
		CbteNro  int64 `json:"CbteNro"`

		Concepto int32   `json:"Concepto"`
		DocTipo  int32   `json:"DocTipo"`
		DocNro   int64   `json:"DocNro"`
		MonId    string  `json:"MonId"`
		MonCotiz float64 `json:"MonCotiz"`

		ImpTotal   float64 `json:"ImpTotal"`
		ImpTotConc float64 `json:"ImpTotConc"`
		ImpNeto    float64 `json:"ImpNeto"`
		ImpOpEx    float64 `json:"ImpOpEx"`
		ImpTrib    float64 `json:"ImpTrib"`
		ImpIVA     float64 `json:"ImpIVA"`

		FchVtoPago   *string `json:"FchVtoPago,omitempty"`
		FchServDesde *string `json:"FchServDesde,omitempty"`
		FchServHasta *string `json:"FchServHasta,omitempty"`

		IVAs        []InvoiceTax               `json:"Iva,omitempty"`
		Periodo     *InvoicePeriod             `json:"Periodo,omitempty"`
		Tributos    []InvoiceTribute           `json:"Tributos,omitempty"`
		CbtesAsoc   []InvoiceAssociatedVoucher `json:"CbtesAsoc,omitempty"`
		Opcionales  []InvoiceOptional          `json:"Opcionales,omitempty"`
		Compradores []InvoiceBuyer             `json:"Compradores,omitempty"`
		Actividades []InvoiceActivity          `json:"Actividades,omitempty"`
	}

	GetInvoiceDataResponse struct {
		PtoVta          int32  `json:"PtoVta,omitempty"`
		Resultado       string `json:"Resultado,omitempty"`
		EmisionTipo     string `json:"EmisionTipo,omitempty"`
		CodAutorizacion string `json:"CodAutorizacion,omitempty"`

		DocTipo int32 `json:"DocTipo,omitempty"`
		DocNro  int64 `json:"DocNro,omitempty"`

		Concepto  int32  `json:"Concepto,omitempty"`
		CbteTipo  int32  `json:"CbteTipo,omitempty"`
		CbteDesde int64  `json:"CbteDesde,omitempty"`
		CbteHasta int64  `json:"CbteHasta,omitempty"`
		CbteFch   string `json:"CbteFch,omitempty"`

		ImpTotal   float64 `json:"ImpTotal,omitempty"`
		ImpTotConc float64 `json:"ImpTotConc,omitempty"`
		ImpNeto    float64 `json:"ImpNeto,omitempty"`
		ImpOpEx    float64 `json:"ImpOpEx,omitempty"`
		ImpTrib    float64 `json:"ImpTrib,omitempty"`
		ImpIVA     float64 `json:"ImpIVA,omitempty"`

		FchVto       string `json:"FchVto,omitempty"`
		FchProceso   string `json:"FchProceso,omitempty"`
		FchVtoPago   string `json:"FchVtoPago,omitempty"`
		FchServDesde string `json:"FchServDesde,omitempty"`
		FchServHasta string `json:"FchServHasta,omitempty"`

		MonId                  string  `json:"MonId,omitempty"`
		MonCotiz               float64 `json:"MonCotiz,omitempty"`
		CanMisMonExt           string  `json:"CanMisMonExt,omitempty"`
		CondicionIVAReceptorId int32   `json:"CondicionIVAReceptorId,omitempty"`

		Iva           []InvoiceTax               `json:"Iva,omitempty"`
		Tributos      []InvoiceTribute           `json:"Tributos,omitempty"`
		CbtesAsoc     []InvoiceAssociatedVoucher `json:"CbtesAsoc,omitempty"`
		Opcionales    []InvoiceOptional          `json:"Opcionales,omitempty"`
		Compradores   []InvoiceBuyer             `json:"Compradores,omitempty"`
		Actividades   []InvoiceActivity          `json:"Actividades,omitempty"`
		PeriodoAsoc   *InvoicePeriod             `json:"PeriodoAsoc,omitempty"`
		Observaciones []InvoiceObservation       `json:"Observaciones,omitempty"`
	}
)
