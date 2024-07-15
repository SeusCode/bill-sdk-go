package main

import (
	"fmt"
	"os"

	"github.com/seuscode/afip-sdk-go/domain/fiscal"
	"github.com/seuscode/afip-sdk-go/gofip"
	"github.com/seuscode/afip-sdk-go/pkg/billing"
)

func Example2() {
	/*
	 To execute this test you have to replace the following path
	 for a path to your certificate and private key or create the
	 certs folder in the root of this project and put them inside.
	*/
	certificate, err := os.ReadFile("./certs/web.crt")
	if err != nil {
		panic(fmt.Errorf("could not read the certificate %w", err))
	}

	privateKey, err := os.ReadFile("./certs/web.key")
	if err != nil {
		panic(fmt.Errorf("could not read the private key %w", err))
	}

	afip, err := gofip.NewGofip(gofip.GofipOptions{
		Certificate:     certificate,
		PrivateKey:      privateKey,
		TaxId:           20000000000,
		BusinessName:    "BUSINESS NAME",
		FiscalAddress:   "Some address",
		FiscalType:      fiscal.IVA_RESPONSABLE_INSCRIPTO,
		StartOfActivity: "01/10/2000",
		PointOfSale:     1,
		Enviroment:      gofip.PRODUCTION,
	})

	if err != nil {
		panic(err)
	}

	err = afip.GetAuthToken()
	if err != nil {
		panic(fmt.Errorf("error happend on GetAuthToken %w", err))
	}

	afipElectronicBilling := billing.NewElectronicBilling(afip, billing.BillingOptions{})
	resp, err := afipElectronicBilling.GetAliquotTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
