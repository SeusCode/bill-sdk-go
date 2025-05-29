package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/seuscode/bill-sdk-go/v2/api/afip"
	"github.com/seuscode/bill-sdk-go/v2/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/v2/models/afip/document"
	"github.com/seuscode/bill-sdk-go/v2/models/afip/invoice"
	"github.com/seuscode/bill-sdk-go/v2/models/api"
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

	vtoPago := time.Now().Add(10 * 24 * time.Hour).Format("20060102")

	invoice, err := afip.EBilling.IssueInvoice(&invoice.IssueInvoiceRequest{
		CbteTipo:   invoice.FACTURA_C,
		Concepto:   invoice.SERVICIOS,
		DocTipo:    document.CF,
		DocNro:     0,
		FchVtoPago: &vtoPago,
		Items: []invoice.InvoiceItem{
			{
				Id:       "1",
				Cantidad: 1,
				Iva:      aliquot.ZeroPercent,
				Precio:   50,
				Desc:     "Producto de prueba",
				Subtotal: 50,
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
