package afip

import (
	"fmt"
	"strings"

	"github.com/seuscode/bill-sdk-go/models/afip/citizen"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type citizenRegistry interface {
	/**
	 * Asks to AFIP Servers for person information {@see WS
	 * Specification item 4.10}
	 *
	 * @return array with information of the target
	 **/
	GetCitizenRecord(registryNumber int, document string, documentType document.DocumentType) (*citizen.GetCitizenRecordResponse, *http.ApiErrorDetails)
}

type cRegistry struct {
	afip *AfipData
}

func newCitizenRegistry(afip *AfipData) citizenRegistry {
	return &cRegistry{
		afip: afip,
	}
}

/*
	-=============================-
	-== Functions Implementation =-
	-=============================-
*/

func (c *cRegistry) GetCitizenRecord(registryNumber int, citizenDocument string, documentType document.DocumentType) (*citizen.GetCitizenRecordResponse, *http.ApiErrorDetails) {
	var resp citizen.GetCitizenRecordResponse

	endpoint := citizenDocument
	if registryNumber == 13 {
		if documentType == document.DNI {
			endpoint = fmt.Sprintf("%s/%s", "dni", citizenDocument)
		} else {
			endpoint = fmt.Sprintf("%s/%s", "cuit", citizenDocument)
		}
	}

	err := c.afip.HttpClient.Get(fmt.Sprintf("%s/%s", strings.ReplaceAll(ENDPOINT_CITIZEN_DATA, "{padronId}", fmt.Sprintf("%d", registryNumber)), endpoint), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
