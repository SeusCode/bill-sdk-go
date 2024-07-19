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

	resp, err := afip.EBilling.GetAliquotTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ aliquots ================")
	fmt.Println(resp)
	fmt.Println("================ aliquots end ================")

	conceptResponse, err := afip.EBilling.GetConceptTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ concepts ================")
	fmt.Println(*conceptResponse)
	fmt.Println("================ concepts end ================")

	docResponse, err := afip.EBilling.GetDocumentTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ documents ================")
	fmt.Println(docResponse)
	fmt.Println("================ documents end ================")

	currenciesResponse, err := afip.EBilling.GetCurrenciesTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ currencies ================")
	fmt.Println(currenciesResponse)
	fmt.Println("================ currencies end ================")

	posResponse, err := afip.EBilling.GetSalesPoints()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ pointofsales ================")
	fmt.Println(posResponse)
	fmt.Println("================ pointofsales end ================")

	taxResponse, err := afip.EBilling.GetTaxTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ taxtypes ================")
	fmt.Println(taxResponse)
	fmt.Println("================ taxtypes end ================")

	voucherResponse, err := afip.EBilling.GetVoucherTypes(false)

	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n================ vouchertypes ================")
	fmt.Println(voucherResponse)
	fmt.Println("================ vouchertypes end ================")
}
