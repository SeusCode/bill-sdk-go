package fiscal

type (
	FiscalType string
)

const (
	// Fiscal Types
	IVA_EXENTO                   FiscalType = "IVA Exento"
	IVA_RESPONSABLE_INSCRIPTO    FiscalType = "IVA Responsable Inscripto"
	IVA_RESPONSABLE_NO_INSCRIPTO FiscalType = "IVA Responsable No Inscripto"

	MONOTRIBUTO            FiscalType = "Responsable Monotributo"
	NO_CATEGORIZADO        FiscalType = "Sujeto No Categorizado"
	CONSUMIDOR_FINAL       FiscalType = "Consumidor Final"
	PROVEEDOR_DEL_EXTERIOR FiscalType = "Proveedor del Exterior"
)
