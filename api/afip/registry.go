package afip

import (
	"errors"
	"fmt"
	"strings"

	"github.com/seuscode/bill-sdk-go/models/afip/citizen"
	"github.com/seuscode/bill-sdk-go/models/afip/document"
	"github.com/seuscode/bill-sdk-go/pkg/endpoints"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type citizenRegistry interface {
	/**
	 * Asks to AFIP Servers for person information {@see WS
	 * Specification item 4.10}
	 *
	 * @return array with information of the target
	 **/
	GetPersonInformation(registryNumber int, document string, documentType document.DocumentType) (*citizen.GetPersonInformationResponse, error)
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

func (c *cRegistry) GetPersonInformation(registryNumber int, citizenDocument string, documentType document.DocumentType) (*citizen.GetPersonInformationResponse, error) {
	var resp citizen.GetPersonInformationResponse

	if documentType != document.CUIT && documentType != document.DNI && documentType != document.CUIL {
		return nil, errors.New("document type not allowed")
	}

	endpoint := citizenDocument
	if registryNumber == 13 {
		if documentType == document.DNI {
			endpoint = fmt.Sprintf("%s/%s", "dni", citizenDocument)
		} else {
			endpoint = fmt.Sprintf("%s/%s", "cuit", citizenDocument)
		}
	}

	apiData, err := c.afip.HttpClient.Get(fmt.Sprintf("%s/%s", strings.ReplaceAll(endpoints.CITIZEN_DATA, "{padronId}", fmt.Sprintf("%d", registryNumber)), endpoint), &resp)
	if err != nil {
		return nil, err
	}

	if apiData.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiData.Status.Code, apiData.Status.Description)
	}

	return &resp, nil
}
