package main

import (
	"encoding/json"
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/invoice"
	"github.com/seuscode/bill-sdk-go/models/api"
)

func main() {
	afip, err := afip.NewAfipManager(afip.AfipOptions{
		ApiKey:     "YOUR_API_KEY",
		Enviroment: api.PRODUCTION,
		Language:   api.ENGLISH,
	})

	if err != nil {
		panic(fmt.Errorf("error creating afip manager: %w", err))
	}

	// Replace parameters below with correct data
	// In this case, we are going to fetch the first invoice of type 1 (Factura A) made in the point of sale 2
	res, infoErr := afip.EBilling.GetIssuedInvoiceData(2, 1, invoice.FACTURA_A)
	if infoErr != nil {
		err, _ := json.Marshal(infoErr)
		fmt.Println(string(err))
		return
	}

	jsonRes, _ := json.Marshal(res)
	fmt.Println(string(jsonRes))
}
