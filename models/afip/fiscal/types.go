package fiscal

type (
	FiscalType uint
)

const (
	// Fiscal Types
	MONOTRIBUTO                  FiscalType = 1
	NO_CATEGORIZADO              FiscalType = 2
	CONSUMIDOR_FINAL             FiscalType = 3
	PROVEEDOR_DEL_EXTERIOR       FiscalType = 4
	IVA_EXENTO                   FiscalType = 5
	IVA_RESPONSABLE_INSCRIPTO    FiscalType = 6
	IVA_RESPONSABLE_NO_INSCRIPTO FiscalType = 7
)
