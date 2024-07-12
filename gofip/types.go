package gofip

import "github.com/seuscode/afip-sdk-go/domain/fiscal"

type (
	Enviroment uint

	GofipOptions struct {
		// Certificado
		Certificate []byte

		// llave privada del certificado
		PrivateKey []byte

		// (por omision o igual a null se establece igual a $cuit) Ingresos Brutos del contribuyente (necesario para elaborar la factura en PDF)
		IIBB *string

		// CUIT del contribuyente (11bytes) (necesario para elaborar la factura en PDF y para factura electronica de AFIP)
		TaxId int

		// Razon social/Nombre y apellido del contribuyente (necesario para elaborar la factura en PDF y para factura electronica de AFIP)
		BusinessName string

		//  Marca o Nombre de fantasía del contribuyente (puede ser vacío, pero no nulo, en cuyo caso se completará con el nombre del contribuyente $razonSocial)
		ComercialName *string

		// Domicilio comercial del contribuyente (necesario para elaborar la factura en PDF)
		FiscalAddress string

		// IVA Responsable Inscripto / IVA Responsable No Inscripto / IVA Exento / Consumidor Final / Responsable Monotributo / Sujeto No Categorizado / Proveedor del Exterior  (necesario para elaborar la factura en PDF)
		FiscalType fiscal.FiscalType

		// Inicio de actividades del contribuyente (dd/mm/aaaa) (necesario para elaborar la factura en PDF)
		StartOfActivity string

		// Punto de venta (necesario para elaborar la factura en PDF y para factura electronica de AFIP)
		PointOfSale int

		// Entorno de ejecucion
		Enviroment Enviroment
	}
)

const (
	// Enviroment Types
	TESTING    Enviroment = 1
	PRODUCTION Enviroment = 0
)
