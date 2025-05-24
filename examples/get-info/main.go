package main

import (
	"encoding/json"
	"fmt"

	"github.com/seuscode/bill-sdk-go/v2/api/afip"
	"github.com/seuscode/bill-sdk-go/v2/models/api"
)

func main() {
	afip, initErr := afip.NewAfipManager(afip.AfipOptions{
		ApiKey:     "YOUR_API_KEY",
		Enviroment: api.PRODUCTION,
		Language:   api.SPANISH,
	})

	if initErr != nil {
		panic(initErr)
	}

	fmt.Println("\n\n================ aliquots ================")
	resp, err := afip.EBilling.GetAliquots()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ := json.Marshal(resp)
	fmt.Println(string(jsonResp))
	fmt.Println("================ aliquots end ================")

	fmt.Println("\n\n================ concepts ================")
	conceptResponse, err := afip.EBilling.GetConcepts()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(*conceptResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ concepts end ================")

	fmt.Println("\n\n================ documents ================")
	docResponse, err := afip.EBilling.GetDocuments()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(docResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ documents end ================")

	fmt.Println("\n\n================ currencies ================")
	currenciesResponse, err := afip.EBilling.GetCurrencies()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(currenciesResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ currencies end ================")

	fmt.Println("\n\n================ pointofsales ================")
	posResponse, err := afip.EBilling.GetSalesPoints()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(posResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ pointofsales end ================")

	fmt.Println("\n\n================ tribute types ================")
	taxResponse, err := afip.EBilling.GetTributes()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(taxResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ tribute types end ================")

	fmt.Println("\n\n================ vouchertypes ================")
	voucherResponse, err := afip.EBilling.GetVouchers()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(voucherResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ vouchertypes end ================")

	fmt.Println("\n\n================ payment methods ================")
	pmResponse, err := afip.EBilling.GetPaymentMethods()

	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(pmResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ payment methods end ================")

	fmt.Println("\n\n================ optionals ================")
	opResponse, err := afip.EBilling.GetOptionals()
	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(opResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ optionals end ================")

	fmt.Println("\n\n================ currency exchange ================")
	cotResponse, err := afip.EBilling.GetCurrencyExchangeRate("DOL")
	if err != nil {
		parsedErr, _ := json.Marshal(err)
		panic(string(parsedErr))
	}

	jsonResp, _ = json.Marshal(cotResponse)
	fmt.Println(string(jsonResp))
	fmt.Println("================ currency exchange end ================")
}
