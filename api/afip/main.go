package afip

import (
	"errors"
	"strings"

	"github.com/seuscode/bill-sdk-go/v2/config"
	"github.com/seuscode/bill-sdk-go/v2/models/api"
	"github.com/seuscode/bill-sdk-go/v2/pkg/http"
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

		// Language to receive response error messages
		Language api.Language
	}
)

func NewAfipManager(opts AfipOptions) (*AfipData, error) {
	if opts.ApiKey == "" || (opts.Enviroment != api.TESTING && opts.Enviroment != api.PRODUCTION) {
		return nil, errors.New("missing a required parameters")
	}

	if opts.Language != api.SPANISH && opts.Language != api.ENGLISH {
		return nil, errors.New("language not allowed")
	}

	// Set the base URL based on the environment
	afipManager := &AfipData{
		apiKey:     opts.ApiKey,
		enviroment: opts.Enviroment,
	}

	// Set the base URL based on the environment
	baseURL := config.API_BASE_URL
	switch opts.Enviroment {
	case api.PRODUCTION:
		baseURL = strings.ReplaceAll(baseURL, "{env}", "production")
	case api.TESTING:
		baseURL = strings.ReplaceAll(baseURL, "{env}", "test")
	default:
		return nil, errors.New("invalid environment")
	}

	// Set the base URL for the API
	afipManager.HttpClient = http.NewHttpClient(&afipManager.apiKey, baseURL, opts.Language)
	afipManager.EBilling = newElectronicBilling(afipManager)
	afipManager.Registry = newCitizenRegistry(afipManager)

	return afipManager, nil
}

func (g *AfipData) ServerPing() (*PingResponse, *http.ApiErrorDetails) {
	var res PingResponse

	err := g.HttpClient.Get(ENDPOINT_PING, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (g *AfipData) AfipServerStatus() (*ServerStatusResponse, *http.ApiErrorDetails) {
	var res ServerStatusResponse

	err := g.HttpClient.Get(ENDPOINT_AFIP_STATUS, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
