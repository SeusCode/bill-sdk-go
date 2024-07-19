package afip

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/seuscode/bill-sdk-go/config"
	"github.com/seuscode/bill-sdk-go/models/afip/auth"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
	"github.com/seuscode/bill-sdk-go/models/api"
	"github.com/seuscode/bill-sdk-go/pkg/endpoints"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type (
	afipData struct {
		// Certificate
		certificate []byte

		// Certificate private key
		privateKey []byte

		// taxpayer gross income
		IIBB *string

		// taxpayer single tax identification (CUIT)
		TaxId int

		// businessName / Name & surname of the taxpayer
		BusinessName string

		//  Comercial name to be used on invoices instead of businessName
		ComercialName *string

		// Address registered inside afip
		FiscalAddress string

		// fiscal type of the taxpayer
		FiscalType fiscal.FiscalType

		// Taxpayer activity start date
		StartOfActivity string

		// Point of sale that will be used on invoices
		PointOfSale int

		// Api enviroment to make the requests
		enviroment api.Enviroment

		// Authorization token to access api
		authToken string

		// Http client to make the requests
		HttpClient *http.HttpClient

		// Electronic billing interface to access the functions
		EBilling electronicBilling
	}

	AfipOptions struct {
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
		Enviroment api.Enviroment
	}
)

func NewAfipManager(opts AfipOptions) (*afipData, error) {
	if opts.TaxId == 0 || (opts.Enviroment != api.TESTING && opts.Enviroment != api.PRODUCTION) {
		return nil, errors.New("missing a required option")
	}

	afipManager := &afipData{
		certificate: opts.Certificate,
		privateKey:  opts.PrivateKey,

		TaxId:           opts.TaxId,
		IIBB:            opts.IIBB,
		BusinessName:    opts.BusinessName,
		ComercialName:   opts.ComercialName,
		FiscalAddress:   opts.FiscalAddress,
		FiscalType:      opts.FiscalType,
		StartOfActivity: opts.StartOfActivity,
		PointOfSale:     opts.PointOfSale,

		enviroment: opts.Enviroment,
	}

	afipManager.HttpClient = http.NewHttpClient(&afipManager.authToken, config.API_BASE_URL)
	afipManager.EBilling = newElectronicBilling(afipManager)

	return afipManager, nil
}

func (g *afipData) GetAuthToken() error {
	r := auth.AuthRequest{
		Pos:          g.PointOfSale,
		TaxId:        g.TaxId,
		BusinessName: g.BusinessName,

		FiscalType:      g.FiscalType,
		FiscalAddress:   g.FiscalAddress,
		StartOfActivity: g.StartOfActivity,

		Enviroment: int(g.enviroment),
	}

	if g.ComercialName != nil {
		r.ComercialName = *g.ComercialName
	} else {
		r.ComercialName = ""
	}

	if g.IIBB != nil {
		r.IIBB = *g.IIBB
	} else {
		r.IIBB = fmt.Sprintf("%d", g.TaxId)
	}

	r.Certificate = base64.StdEncoding.EncodeToString(g.certificate)
	r.CertificateKey = base64.StdEncoding.EncodeToString(g.privateKey)

	var authResponse auth.AuthResponse

	if respStatus, err := g.HttpClient.Post(endpoints.AUTH, r, &authResponse); err != nil || respStatus.Status.Type != http.SUCCESS {
		if err != nil {
			return err
		}

		return fmt.Errorf("[ERROR] (%s): %s", respStatus.Status.Code, respStatus.Status.Description)
	}

	g.authToken = authResponse.JWT

	fmt.Println("[SUCCESS] (Login): Access approved by afip wsaa")
	return nil
}

func (g *afipData) ServerPing() error {
	return nil
}

func (g *afipData) SessionAlive() error {
	return nil
}
