package afip

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/seuscode/afip-sdk-go/domain/api/requests"
	"github.com/seuscode/afip-sdk-go/domain/api/responses"
	"github.com/seuscode/afip-sdk-go/domain/fiscal"
	"github.com/seuscode/afip-sdk-go/endpoints"
	"github.com/seuscode/afip-sdk-go/pkg/billing"
)

type Gofip struct {
	// SDK version
	version string

	// Certificate
	certificate []byte

	// Certificate private key
	privateKey []byte

	// taxpayer gross income
	iibb *string

	// taxpayer single tax identification (CUIT)
	taxId string

	// businessName / Name & surname of the taxpayer
	businessName string

	//  Comercial name to be used on invoices instead of businessName
	comercialName *string

	// Address registered inside afip
	fiscalAddress string

	// fiscal type of the taxpayer
	fiscalType fiscal.FiscalType

	// Taxpayer activity start date
	startOfActivity string

	// Point of sale that will be used on invoices
	pointOfSale int

	// Api enviroment to make the requests
	enviroment Enviroment

	// Authorization token to access api
	authToken string

	// Http client to make the requests
	httpClient *httpClient

	ElectronicBilling *billing.ElectronicBilling
}

func NewGofip(opts GofipOptions) (*Gofip, error) {
	if opts.TaxId == "" || (opts.Enviroment != TESTING && opts.Enviroment != PRODUCTION) {
		return nil, errors.New("missing a required option")
	}

	gofip := &Gofip{
		version:     "1.0.0",
		certificate: opts.Certificate,
		privateKey:  opts.PrivateKey,

		taxId:           opts.TaxId,
		iibb:            opts.IIBB,
		businessName:    opts.BusinessName,
		comercialName:   opts.ComercialName,
		fiscalAddress:   opts.FiscalAddress,
		fiscalType:      opts.FiscalType,
		startOfActivity: opts.StartOfActivity,
		pointOfSale:     opts.PointOfSale,

		enviroment: opts.Enviroment,
	}

	gofip.httpClient = newHttpClient(gofip)
	if err := gofip.GetAuthToken(); err != nil {
		fmt.Println("[WARNING] (LOGIN): Cannot obtain a jwt")
	}

	return gofip, nil
}

func (g *Gofip) GetAuthToken() error {
	tempHttpClient := newHttpClient(g)

	r := requests.AuthRequest{
		Pos:          g.pointOfSale,
		TaxId:        g.taxId,
		BusinessName: g.businessName,

		FiscalType:      g.fiscalType,
		FiscalAddress:   g.fiscalAddress,
		StartOfActivity: g.startOfActivity,

		Enviroment: int(g.enviroment),
	}

	if g.comercialName != nil {
		r.ComercialName = *g.comercialName
	}

	if g.iibb != nil {
		r.IIBB = *g.iibb
	}

	r.Certificate = base64.StdEncoding.EncodeToString(g.certificate)
	r.CertificateKey = base64.StdEncoding.EncodeToString(g.privateKey)

	var authResponse responses.AuthResponse

	if respStatus, err := tempHttpClient.Post(endpoints.AUTH, r, &authResponse); err != nil || respStatus.Status.Type != SUCCESS {
		if err != nil {
			return err
		}

		return fmt.Errorf("[ERROR] (%s): %s", respStatus.Status.Code, respStatus.Status.Description)
	}

	g.authToken = authResponse.JWT
	g.httpClient = newHttpClient(g)

	return nil
}

func (g *Gofip) ServerPing() error {
	return nil
}

func (g *Gofip) SessionAlive() error {
	return nil
}
