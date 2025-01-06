package main

import (
	"encoding/json"
	"fmt"

	"github.com/seuscode/bill-sdk-go/api/afip"
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

	// Replace parameters below with correct data
	res, err := afip.EBilling.GetVoucherInfo(5, 7, voucher.FacturaA)
	if err != nil {
		panic(err)
	}

	jsonData, _ := json.Marshal(res)
	fmt.Println(string(jsonData))
}
