package main

import (
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/aliquot"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
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

	var resp voucher.CreateVoucherResponse
	err = afip.EBilling.CreateVoucher(&voucher.Voucher{
		CbteTipo:     voucher.FacturaB,
		Concepto:     voucher.Productos,
		MetodoDePago: payment_method.Cash,
		DocTipo:      document.CF,
		DocNro:       0,

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
	}, &resp)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
