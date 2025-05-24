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

	"github.com/seuscode/bill-sdk-go/v2/models/api"
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

	ApiErrorDetails struct {
		Type     string `json:"type,omitempty"`
		Title    string `json:"title"`
		Status   int    `json:"status"`
		Detail   string `json:"detail"`
		Instance string `json:"instance,omitempty"`
		Code     string `json:"code,omitempty"`
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

func NewHttpClient(authToken *string, baseURL string, lang api.Language) *HttpClient {
	client := &http.Client{
		Timeout: 90 * time.Second,
	}

	// Adding default headers
	defaultHeaders := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Language": string(lang),
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

func (*HttpClient) HandleErrorResponse(resp *http.Response) *ApiErrorDetails {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var errDetails ApiErrorDetails
	err := json.Unmarshal(bodyBytes, &errDetails)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling server response into error details: %v", err))
	}

	return &errDetails
}

func (c *HttpClient) Get(endpoint string, response interface{}) *ApiErrorDetails {
	c.validateToken()

	req, err := http.NewRequest(GET, fmt.Sprintf("%s%s", c.BaseURL, endpoint), nil)
	if err != nil {
		panic(fmt.Errorf("error creating request: %v", err))
	}

	resp, err := c.http.Do(req)
	if err != nil {
		panic(fmt.Errorf("error doing request: %v", err))
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		return err
	}

	if err := c.bodyParser(resp, response); err != nil {
		panic(fmt.Errorf("error parsing response body: %v", err))
	}

	return nil
}

func (c *HttpClient) Post(endpoint string, data interface{}, expectedResponse interface{}) *ApiErrorDetails {
	c.validateToken()

	type apiRequest struct {
		Data interface{} `json:"data"`
	}

	apiReq := apiRequest{
		Data: data,
	}

	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		panic(fmt.Errorf("error marshaling request data: %v", err))
	}

	req, err := http.NewRequest(POST, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		panic(fmt.Errorf("error creating request: %v", err))
	}

	resp, err := c.http.Do(req)
	if err != nil {
		panic(fmt.Errorf("error doing request: %v", err))
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		return err
	}

	if err := c.bodyParser(resp, expectedResponse); err != nil {
		panic(fmt.Errorf("error parsing response body: %v", err))
	}

	return nil
}

func (c *HttpClient) PostWithFileOnResponse(endpoint string, data interface{}, folderName, fileName string) (string, error) {
	c.validateToken()

	type apiRequest struct {
		Data interface{} `json:"data"`
	}

	apiReq := apiRequest{
		Data: data,
	}

	jsonData, err := json.Marshal(apiReq)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(POST, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}

	if err := c.HandleErrorResponse(resp); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("%v: %s", err, string(bodyBytes))
	}

	return c.saveResponseFile(resp, folderName, fileName)
}

func (c *HttpClient) bodyParser(r *http.Response, expectedResponse interface{}) error {
	// Read response body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// If not response expected, return nil
	if expectedResponse == nil {
		return nil
	}

	// Check if expectedResponse is a pointer
	if reflect.ValueOf(expectedResponse).Kind() != reflect.Ptr {
		panic(fmt.Errorf("expected %s as expectedResponse, %s received", reflect.Ptr, reflect.ValueOf(expectedResponse).Kind()))
	}

	err = json.Unmarshal(body, expectedResponse)
	if err != nil {
		return err
	}

	return nil
}

func (c *HttpClient) saveResponseFile(r *http.Response, folderName, fileName string) (string, error) {
	// Create or truncate the file to save the PDF
	finalPath := path.Clean(fmt.Sprintf("./%s/%s", folderName, fileName))

	err := os.MkdirAll(path.Dir(finalPath), 0777)
	if err != nil {
		return "", fmt.Errorf("error creating folder: %v", err)
	}

	file, err := os.Create(finalPath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Copy the response body content to the file
	_, err = io.Copy(file, r.Body)
	if err != nil {
		return "", fmt.Errorf("error copying response body to file: %v", err)
	}

	return finalPath, nil
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
