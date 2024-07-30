package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

type (
	ApiStatusType string

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

	HttpClient struct {
		http      *http.Client
		BaseURL   string
		headers   map[string]string
		authToken *string
	}
)

const (
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

func NewHttpClient(authToken *string, baseURL string) *HttpClient {
	client := &http.Client{
		Timeout: 90 * time.Second,
	}

	// Adding default headers
	defaultHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if authToken != nil {
		defaultHeaders["Authorization"] = "Bearer " + *authToken
	}

	client.Transport = &transportWithHeaders{
		headers: defaultHeaders,
		rt:      http.DefaultTransport,
	}

	return &HttpClient{
		http:      client,
		BaseURL:   baseURL,
		headers:   defaultHeaders,
		authToken: authToken,
	}
}

func (*HttpClient) HandleErrorResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("error: status code %d, status: %s", resp.StatusCode, resp.Status)
}

func (c *HttpClient) Get(endpoint string, expectedResponse interface{}) (*ApiResponse, error) {
	c.validateToken()

	req, err := http.NewRequest(GET, fmt.Sprintf("%s%s", c.BaseURL, endpoint), nil)
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

func (c *HttpClient) Post(endpoint string, data interface{}, expectedResponse interface{}) (*ApiResponse, error) {
	c.validateToken()

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

	req, err := http.NewRequest(POST, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
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

func (c *HttpClient) PostWithFileOnResponse(endpoint string, data interface{}, folderName, fileName string) error {
	c.validateToken()

	type apiRequest struct {
		Data interface{} `json:"data"`
	}

	apiReq := apiRequest{
		Data: data,
	}

	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(POST, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%w: %s", err, string(bodyBytes))
	}

	if err := c.saveResponseFile(resp, folderName, fileName); err != nil {
		return err
	}

	return nil
}

func (c *HttpClient) bodyParser(r *http.Response, expectedResponse interface{}) (*ApiResponse, error) {
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

func (c *HttpClient) saveResponseFile(r *http.Response, folderName, fileName string) error {
	// Create or truncate the file to save the PDF
	finalPath := path.Clean(fmt.Sprintf("./%s/%s", folderName, fileName))

	err := os.MkdirAll(path.Dir(finalPath), 0777)
	if err != nil {
		return fmt.Errorf("error creating folder: %v", err)
	}

	file, err := os.Create(finalPath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Copy the response body content to the file
	_, err = io.Copy(file, r.Body)
	if err != nil {
		return fmt.Errorf("error copying response body to file: %v", err)
	}

	return nil
}

func (c *HttpClient) validateToken() {
	if c.authToken == nil || strings.Contains(c.headers["Authorization"], *c.authToken) {
		return
	}

	c.headers["Authorization"] = "Bearer " + *c.authToken
	c.http.Transport = &transportWithHeaders{
		headers: c.headers,
		rt:      http.DefaultTransport,
	}

}
