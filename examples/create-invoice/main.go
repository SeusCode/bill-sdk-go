package main

import (
	"encoding/json"
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/invoice"
	"github.com/seuscode/bill-sdk-go/models/api"
)

func main() {
	afip, initErr := afip.NewAfipManager(afip.AfipOptions{
		ApiKey:     "YOUR_API_KEY",
		Enviroment: api.PRODUCTION,
		Language:   api.ENGLISH,
	})

	if initErr != nil {
		panic(initErr)
	}

	invoice, err := afip.EBilling.IssueInvoice(&invoice.IssueInvoiceRequest{
		CbteTipo: invoice.FACTURA_B,
		Concepto: invoice.PRODUCTOS,
		DocTipo:  document.CF,
		DocNro:   0,
		Items: []invoice.InvoiceItem{
			{
				Id:       "124",
				Cantidad: 4.3,
				Iva:      aliquot.TenDotFivePercent,
				Precio:   50,
				Desc:     "Vacio",
				Subtotal: 215,
			},
			{
				Id:       "99124",
				Cantidad: 1,
				Iva:      aliquot.TenDotFivePercent,
				Precio:   -200,
				Desc:     "Descuento en vacio",
				Subtotal: -200,
			},
		},
	})

	if err != nil {
		errJson, _ := json.Marshal(err)
		fmt.Println(string(errJson))
		return
	}

	invJson, _ := json.Marshal(invoice)
	fmt.Println(string(invJson))
}
