package requests

import "github.com/seuscode/bill-sdk-go/domain/fiscal"

type (
	AuthRequest struct {
		Pos int `json:"puntoVenta"`

		BusinessName  string `json:"razonSocial"`
		ComercialName string `json:"nombreComercial"`

		IIBB            string            `json:"ingresosBrutos"`
		TaxId           int               `json:"cuit"`
		FiscalType      fiscal.FiscalType `json:"iva"`
		FiscalAddress   string            `json:"domicilio"`
		StartOfActivity string            `json:"inicioActividad"`

		Certificate    string `json:"certificado"`
		CertificateKey string `json:"certificadoKey"`
		Enviroment     int    `json:"debugMode"`
	}
)
