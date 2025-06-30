package main

import (
	"fmt"

	"github.com/seuscode/bill-sdk-go/v2/api/afip"
	"github.com/seuscode/bill-sdk-go/v2/models/afip/invoice"
	"github.com/seuscode/bill-sdk-go/v2/models/afip/receipt"
	"github.com/seuscode/bill-sdk-go/v2/models/api"
)

func main() {
	afip, err := afip.NewAfipManager(afip.AfipOptions{
		ApiKey:     "YOUR_API_KEY",
		Enviroment: api.PRODUCTION,
		Language:   api.ENGLISH,
	})

	if err != nil {
		panic(err)
	}

	filePath, err := afip.EBilling.GenerateReceiptPDF(receipt.GenerateReceiptPDFRequest{
		ReceiptMetaData: receipt.ReceiptMetaData{
			ClientDocument:           0,
			ClientPhone:              "-",
			ClientEmail:              "-",
			ClientAddress:            "-",
			ClientFiscalType:         uint(invoice.CONSUMIDOR_FINAL),
			ClientNameOrBusinessName: "Your Client Name",
		},
		ReceiptData: receipt.Receipt{
			ReceiptNbr:  "022-0002051",
			ReceiptDate: "20250625",

			Total:    572000,
			Subtotal: 572000,
		},
		ReceiptProducts: []receipt.ReceiptItem{
			{
				Id:       "1",
				Cantidad: 1,
				Price:    572000,
				Desc:     "Desarrollo de Sistema de Facturación y Administración",
				Total:    572000,
			},
		},
	}, "examples/generate-receipt-pdf/receipt", "test_receipt.pdf")

	if err != nil {
		panic(err)
	}

	fmt.Println(filePath)
}
