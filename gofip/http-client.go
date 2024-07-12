package gofip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

type (
	ApiStatusType string

	httpClient struct {
		http *http.Client
	}

	transportWithHeaders struct {
		headers map[string]string
		rt      http.RoundTripper
	}

	ApiStausResponse struct {
		Type        ApiStatusType `json:"type"`
		Code        string        `json:"code"`
		Description string        `json:"description"`
	}

	ApiResponse struct {
		Data   interface{}      `json:"data"`
		Status ApiStausResponse `json:"status"`
	}
)

const (
	baseURL = "http://dev.seuscode.com/afip"

	// Api Status type
	ERROR   ApiStatusType = "error"
	SUCCESS ApiStatusType = "success"

	// Http Methods
	GET  = "GET"
	POST = "POST"
)

func (t *transportWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}
	return t.rt.RoundTrip(req)
}

func newHttpClient(data *Gofip) *httpClient {
	client := &http.Client{
		Timeout: 90 * time.Second,
	}

	// Adding default headers
	defaultHeaders := map[string]string{}

	if data.authToken != "" {
		defaultHeaders["Authorization"] = "Bearer " + data.authToken
	}

	client.Transport = &transportWithHeaders{
		headers: defaultHeaders,
		rt:      http.DefaultTransport,
	}

	return &httpClient{
		http: client,
	}
}

func (*httpClient) HandleErrorResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("error: status code %d, status: %s", resp.StatusCode, resp.Status)
}

func (c *httpClient) Get(endpoint string, expectedResponse interface{}) (*ApiResponse, error) {
	req, err := http.NewRequest(GET, fmt.Sprintf("%s%s", baseURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%w: %s", err, string(bodyBytes))
	}

	apiResponse, err := c.bodyParser(resp, expectedResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}

func (c *httpClient) Post(endpoint string, data interface{}, expectedResponse interface{}) (*ApiResponse, error) {
	type apiRequest struct {
		Data interface{} `json:"data"`
	}

	apiReq := apiRequest{
		Data: data,
	}

	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(POST, fmt.Sprintf("%s%s", baseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%w: %s", err, string(bodyBytes))
	}

	apiResponse, err := c.bodyParser(resp, expectedResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}

func (c *httpClient) bodyParser(r *http.Response, expectedResponse interface{}) (*ApiResponse, error) {
	// Read response body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var resp ApiResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if expectedResponse == nil {
		return &resp, nil
	}

	if reflect.ValueOf(expectedResponse).Kind() != reflect.Ptr {
		panic(fmt.Errorf("expected %s as expectedResponse, %s received", reflect.Ptr, reflect.ValueOf(expectedResponse).Kind()))
	}

	jsonParsedData, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonParsedData, expectedResponse)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}
