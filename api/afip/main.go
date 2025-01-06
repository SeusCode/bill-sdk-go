package afip

import (
	"errors"
	"fmt"

	"github.com/seuscode/bill-sdk-go/config"
	"github.com/seuscode/bill-sdk-go/models/afip/auth"
	"github.com/seuscode/bill-sdk-go/models/api"
	"github.com/seuscode/bill-sdk-go/pkg/endpoints"
	"github.com/seuscode/bill-sdk-go/pkg/http"
)

type (
	AfipData struct {
		// Api enviroment to make the requests
		enviroment api.Enviroment

		// Authorization token to access api
		apiKey string

		// Http client to make the requests
		HttpClient *http.HttpClient

		// Electronic billing interface to access the functions
		EBilling electronicBilling

		// Citizen registry interface to access the functions
		Registry citizenRegistry
	}

	AfipOptions struct {
		// llave privada del api
		ApiKey string

		// Entorno de ejecucion
		Enviroment api.Enviroment
	}
)

type (
	ServerStatusResponse struct {
		ServerStatus auth.ServerStatus `json:"server_status"`
	}

	PingResponse struct {
		Datetime  string `json:"datetime"`
		Timestamp int64  `json:"timestamp"`
	}
)

func NewAfipManager(opts AfipOptions) (*AfipData, error) {
	if opts.ApiKey == "" || (opts.Enviroment != api.TESTING && opts.Enviroment != api.PRODUCTION) {
		return nil, errors.New("missing a required parameters")
	}

	afipManager := &AfipData{
		apiKey:     opts.ApiKey,
		enviroment: opts.Enviroment,
	}

	afipManager.HttpClient = http.NewHttpClient(&afipManager.apiKey, config.API_BASE_URL)
	afipManager.EBilling = newElectronicBilling(afipManager)
	afipManager.Registry = newCitizenRegistry(afipManager)

	return afipManager, nil
}

func (g *AfipData) ServerPing() (*PingResponse, error) {
	var res PingResponse

	apiResponse, err := g.HttpClient.Get(endpoints.PING, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}

func (g *AfipData) AfipServerStatus() (*ServerStatusResponse, error) {
	var res ServerStatusResponse

	apiResponse, err := g.HttpClient.Get(endpoints.AFIP_STATUS, &res)
	if err != nil {
		return nil, err
	}

	if apiResponse.Status.Type != http.SUCCESS {
		return nil, fmt.Errorf("error (%s): %s", apiResponse.Status.Code, apiResponse.Status.Description)
	}

	return &res, nil
}
