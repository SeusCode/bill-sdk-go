package main

import (
	"fmt"
	"os"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
	"github.com/seuscode/bill-sdk-go/models/afip/payment_method"
	"github.com/seuscode/bill-sdk-go/models/afip/voucher"
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

	filePath, err := afip.EBilling.CreatePDF(voucher.VoucherPDF{
		Logo:      "none",
		Watermark: "none",
		Template:  voucher.Clasico,
		PtoVta:    7,

		CbteTipo: voucher.FacturaA,
		CbteNro:  1233,
		CbteFch:  20240730,
		Concepto: voucher.Productos,

		CopiaDesde: 0,
		CopiaHasta: 0,

		DocTipo: document.CUIT,
		DocNro:  20111111112,

		MonId:    "PES",
		MonCotiz: 1,

		CAE:       "6621462155158814",
		CAEFchVto: "20240731",

		Cliente: voucher.VoucherClient{
			Phone:              "39843798543",
			Address:            "Calle 455",
			Email:              "jdoe@test.com",
			SellCondition:      payment_method.Cash,
			FiscalType:         fiscal.CONSUMIDOR_FINAL,
			NameOrBusinessName: "John Doe",
		},

		ImpIVA:     2.1,
		ImpNeto:    10,
		ImpOpEx:    0,
		ImpTotConc: 0,
		ImpTotal:   12.1,
		ImpTrib:    13.1,

		Iva: []voucher.VoucherShare{
			{BaseImp: 10, Id: 5, Total: 2.1},
			{BaseImp: 10, Id: 5, Total: 2.1},
			{BaseImp: 33.7, Id: 4, Total: 12.1},
		},

		Items: []voucher.VoucherItems{
			{
				Id:             "1",
				Qty:            120,
				Price:          1500,
				Discount:       0,
				Desc:           "Producto 1",
				Subtotal:       180000,
				Iva:            aliquot.TwentyOnePercent,
				IvaExento:      false,
				NoGravadodeIVA: false,
			},
		},

		CbtesAsoc: []voucher.AsociatedVouchers{
			{
				Type:  6,
				Pos:   1,
				Nbr:   1,
				TaxId: 20111111112,
			},
		},

		Tributos: []voucher.VoucherTributes{
			{
				Id:      99,
				Desc:    "Ingresos Brutos",
				BaseImp: 150,
				Alic:    5.2,
				Total:   7.8,
			},
		},

		Opcionales: []voucher.VoucherOptionals{
			{
				Id:    17,
				Value: 2,
			},
		},

		Compradores: []voucher.VoucherBuyers{
			{
				DocType:    80,
				DocNro:     20111111112,
				Percentage: 100,
			},
		},
	}, "examples/generate-pdf/invoice", "test_invoice.pdf")

	if err != nil {
		panic(err)
	}

	fmt.Println(filePath)
}
