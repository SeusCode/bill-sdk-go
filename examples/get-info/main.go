package main

import (
	"fmt"
	"os"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
	"github.com/seuscode/bill-sdk-go/models/api"
)

func main() {
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

	afip, err := afip.NewAfipManager(afip.AfipOptions{
		Certificate: certificate,
		PrivateKey:  privateKey,

		// Replace this information with yours
		TaxId:           20123456780,
		BusinessName:    "YOUR BUSINESS NAME",
		FiscalAddress:   "YOUR FISCAL ADDRESS",
		FiscalType:      fiscal.IVA_RESPONSABLE_INSCRIPTO,
		StartOfActivity: "YOUR START OF ACTIVITY DATE", // (dd/mm/yyyy)
		PointOfSale:     1,
		Enviroment:      api.PRODUCTION,
	})

	if err != nil {
		panic(err)
	}

	err = afip.GetAuthToken()
	if err != nil {
		panic(fmt.Errorf("error happend on GetAuthToken %w", err))
	}

	fmt.Println("\n\n================ aliquots ================")
	resp, err := afip.EBilling.GetAliquotTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println("================ aliquots end ================")

	fmt.Println("\n\n================ concepts ================")
	conceptResponse, err := afip.EBilling.GetConceptTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(*conceptResponse)
	fmt.Println("================ concepts end ================")

	fmt.Println("\n\n================ documents ================")
	docResponse, err := afip.EBilling.GetDocumentTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(docResponse)
	fmt.Println("================ documents end ================")

	fmt.Println("\n\n================ currencies ================")
	currenciesResponse, err := afip.EBilling.GetCurrenciesTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(currenciesResponse)
	fmt.Println("================ currencies end ================")

	fmt.Println("\n\n================ pointofsales ================")
	posResponse, err := afip.EBilling.GetSalesPoints()

	if err != nil {
		panic(err)
	}

	fmt.Println(posResponse)
	fmt.Println("================ pointofsales end ================")

	fmt.Println("\n\n================ taxtypes ================")
	taxResponse, err := afip.EBilling.GetTaxTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(taxResponse)
	fmt.Println("================ taxtypes end ================")

	fmt.Println("\n\n================ vouchertypes ================")
	voucherResponse, err := afip.EBilling.GetVoucherTypes(true)

	if err != nil {
		panic(err)
	}

	fmt.Println(voucherResponse)
	fmt.Println("================ vouchertypes end ================")
}
