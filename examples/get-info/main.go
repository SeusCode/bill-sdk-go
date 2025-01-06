package main

import (
	"encoding/json"
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/api"
)

func main() {
	afip, err := afip.NewAfipManager(afip.AfipOptions{
		ApiKey:     "633C3509DC2513BB9E5C414AB542444D6A33F2478C188F4BBB58",
		Enviroment: api.PRODUCTION,
	})

	if err != nil {
		panic(err)
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
	voucherResponse, err := afip.EBilling.GetVoucherTypes()

	if err != nil {
		panic(err)
	}

	fmt.Println(voucherResponse)
	fmt.Println("================ vouchertypes end ================")

	fmt.Println("\n\n================ payment methods ================")
	pmResponse, err := afip.EBilling.GetPaymentMethods()

	if err != nil {
		panic(err)
	}

	js, _ := json.Marshal(pmResponse)
	fmt.Println(string(js))
	fmt.Println("================ payment methods end ================")

	fmt.Println("\n\n================ optionals ================")
	opResponse, err := afip.EBilling.GetOptionalTypes()
	if err != nil {
		panic(err)
	}

	jsO, _ := json.Marshal(opResponse)
	fmt.Println(string(jsO))
	fmt.Println("================ optionals end ================")

	fmt.Println("\n\n================ cotizations ================")
	cotResponse, err := afip.EBilling.GetCurrencyCotization("DOL")
	if err != nil {
		panic(err)
	}

	jsC, _ := json.Marshal(cotResponse)
	fmt.Println(string(jsC))
	fmt.Println("================ cotizations end ================")
}
