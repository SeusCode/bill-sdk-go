package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/seuscode/bill-sdk-go/api/afip"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/models/api"
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

	personIDs := []string{
		// Add here all the taxId of person you want to get the info
	}

	for _, personID := range personIDs {
		res, err := afip.Registry.GetCitizenRecord(13, personID, document.DNI)
		if err != nil {
			log.Fatalf("Error getting information for personID %s\n%v", personID, err.Detail)
		}

		jsonRes, _ := json.Marshal(res)
		fmt.Println(string(jsonRes))
	}
}
