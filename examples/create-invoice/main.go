package main

import (
	"fmt"
	"os"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/afip/fiscal"
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

	var resp voucher.CreateVoucherResponse
	err = afip.EBilling.CreateVoucher(&voucher.Voucher{
		CbteTipo: voucher.FacturaA,
		Concepto: voucher.Productos,

		DocTipo: document.CUIT,
		DocNro:  20335143990,

		Items: []voucher.VoucherItems{
			{
				Id:             "testItem1",
				Qty:            1,
				Iva:            aliquot.TwentyOnePercent,
				Price:          10,
				Desc:           "Test item",
				Discount:       0,
				Subtotal:       10,
				IvaExento:      false,
				NoGravadodeIVA: false,
			},
		},
	}, &resp)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
