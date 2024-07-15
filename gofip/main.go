package gofip

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/seuscode/afip-sdk-go/domain/api/requests"
	"github.com/seuscode/afip-sdk-go/domain/api/responses"
	"github.com/seuscode/afip-sdk-go/domain/fiscal"
	"github.com/seuscode/afip-sdk-go/endpoints"
)

type Gofip struct {
	// SDK version
	version string

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
	enviroment Enviroment

	// Authorization token to access api
	authToken string

	// Http client to make the requests
	HttpClient *httpClient
}

func NewGofip(opts GofipOptions) (*Gofip, error) {
	if opts.TaxId == 0 || (opts.Enviroment != TESTING && opts.Enviroment != PRODUCTION) {
		return nil, errors.New("missing a required option")
	}

	gofip := &Gofip{
		version:     "1.0.0",
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

	gofip.HttpClient = newHttpClient(gofip)

	return gofip, nil
}

func (g *Gofip) GetAuthToken() error {
	tempHttpClient := newHttpClient(g)

	r := requests.AuthRequest{
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

	var authResponse responses.AuthResponse

	if respStatus, err := tempHttpClient.Post(endpoints.AUTH, r, &authResponse); err != nil || respStatus.Status.Type != SUCCESS {
		if err != nil {
			return err
		}

		return fmt.Errorf("[ERROR] (%s): %s", respStatus.Status.Code, respStatus.Status.Description)
	}

	g.authToken = authResponse.JWT
	g.HttpClient = newHttpClient(g)

	fmt.Println("[SUCCESS] (Login): Access approved by afip wsaa")
	return nil
}

func (g *Gofip) ServerPing() error {
	return nil
}

func (g *Gofip) SessionAlive() error {
	return nil
}
