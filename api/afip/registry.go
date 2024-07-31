package afip

import (
	"errors"
	"fmt"

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
	GetPersonInformation(registryNumber int, document int64, documentType document.DocumentType) (*citizen.GetPersonInformationResponse, error)
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

func (c *cRegistry) GetPersonInformation(registryNumber int, citizenDocument int64, documentType document.DocumentType) (*citizen.GetPersonInformationResponse, error) {
	var resp citizen.GetPersonInformationResponse

	if documentType != document.CUIT && documentType != document.DNI {
		return nil, errors.New("document type not allowed")
	}

	r := citizen.GetPersonInformationRequest{
		RegistryNumber: registryNumber,
		TaxId:          citizenDocument,
	}

	if documentType == document.DNI {
		r.TaxId = 0
		r.CitizenId = citizenDocument
	}

	apiData, err := c.afip.HttpClient.Post(endpoints.CITIZEN_DATA, r, &resp)
	if err != nil {
		return nil, err
	}

	if apiData.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiData.Status.Code, apiData.Status.Description)
	}

	return &resp, nil
}
