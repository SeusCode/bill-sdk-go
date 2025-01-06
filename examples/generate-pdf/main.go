package main

import (
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
	"github.com/seuscode/bill-sdk-go/models/afip/voucher"
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

	filePath, err := afip.EBilling.CreatePDF(voucher.CreateVoucherPDFRequest{
		CAE:        "75010562627746",
		CAEFchVto:  "20250115",
		QREndpoint: "https://www.afip.gob.ar/fe/qr/?p=eyJjb2RBdXQiOjc1MDEwNTYyNjI3NzQ2LCJjdHoiOjEsImN1aXQiOjIwMjg5NjgzODEwLCJmZWNoYSI6IjIwMjUtMDEtMDUiLCJpbXBvcnRlIjoxOTEuOTIsIm1vbmVkYSI6IlBFUyIsIm5yb0NtcCI6MjIxMywibnJvRG9",
		AfipInformation: voucher.AFIPInformation{
			VoucherDate:  "20250105",
			VoucherType:  6,
			Concept:      1,
			Document:     0,
			DocumentType: 99,
			ImpIVA:       18.24,
			ImpNeto:      173.68,
			ImpOpEx:      0,
			ImpTotal:     191.92,
			ImpTotConc:   0,
			ImpTrib:      0,
			MonCotiz:     1,
			MonId:        "PES",
			PointOfSale:  7,
			Iva: []voucher.VoucherShare{{
				Id:      4,
				BaseImp: 173.68,
				Total:   18.24,
			}},
			Items: []voucher.VoucherItems{
				{
					Id:             "124",
					Qty:            4.3,
					Iva:            aliquot.TenDotFivePercent,
					Price:          50,
					Desc:           "Vacio",
					Discount:       0,
					Subtotal:       215,
					IvaExento:      false,
					NoGravadodeIVA: false,
				},
				{
					Id:             "0001",
					Qty:            1,
					Iva:            aliquot.TenDotFivePercent,
					Price:          -41.32,
					Desc:           "Descuento",
					Discount:       0,
					Subtotal:       -41.32,
					IvaExento:      false,
					NoGravadodeIVA: false,
				},
			},
		},
		VoucherNbr:               2213,
		MetodoPago:               payment_method.Cash,
		ClientFiscalType:         uint(fiscal.CONSUMIDOR_FINAL),
		ClientPhone:              "+543364123456",
		ClientEmail:              "usuario@gmail.com",
		ClientAddress:            "Alguna calle 235",
		ClientNameOrBusinessName: "algun nombre",
	}, "examples/generate-pdf/invoice", "test_invoice.pdf")

	if err != nil {
		panic(err)
	}

	fmt.Println(filePath)
}
