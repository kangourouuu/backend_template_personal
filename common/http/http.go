package http

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"backend_template_personal/common/log"
)

func PostHeader(url string, header string, body any) (*resty.Response, error) {
	log.Infof("PostHeader | url : %v | headers : %v | body : %v", url, header, body)
	client := resty.New()
	client.SetTimeout(time.Second * 3)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", header).
		SetBody(body).
		Post(url)
}

func Post(url string, body any) (*resty.Response, error) {
	client := resty.New()
	client.SetTimeout(time.Second * 10)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
}

// HTTPClient is a wrapper around the resty client to simplify making HTTP requests.
type HTTPClient struct {
	client *resty.Client
}

// NewHTTPClient creates a new HTTPClient with a configured resty client.
func NewHTTPClient() *HTTPClient {
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})

	return &HTTPClient{client: client}
}

// RequestParams encapsulates the parameters for making an HTTP request.
type RequestParams struct {
	URL     string
	Headers map[string]string
	Body    interface{}
}

// Get executes a GET request and parses the response into the provided struct.
func (h *HTTPClient) Get(url string, headers map[string]string, responseStruct interface{}) error {
	return h.ExecuteRequest("GET", url, headers, nil, responseStruct)
}

// Post executes a POST request and parses the response into the provided struct.
func (h *HTTPClient) Post(url string, headers map[string]string, body interface{}, responseStruct interface{}) error {
	return h.ExecuteRequest("POST", url, headers, body, responseStruct)
}

// Put executes a PUT request and parses the response into the provided struct.
func (h *HTTPClient) Put(url string, headers map[string]string, body interface{}, responseStruct interface{}) error {
	return h.ExecuteRequest("PUT", url, headers, body, responseStruct)
}

// Delete executes a DELETE request and parses the response into the provided struct.
func (h *HTTPClient) Delete(url string, headers map[string]string, body interface{}, responseStruct interface{}) error {
	return h.ExecuteRequest("DELETE", url, headers, body, responseStruct)
}

// ExecuteRequest executes an HTTP request based on the provided parameters.
// ... existing code ...

// ExecuteRequest executes an HTTP request and parses the response into the provided struct.
func (h *HTTPClient) ExecuteRequest(method string, url string, headers map[string]string, body interface{}, responseStruct interface{}) error {
	requestHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	// Merge headers, overwriting defaults with provided headers
	for key, value := range headers {
		requestHeaders[key] = value
	}
	req := h.client.R().
		SetHeaders(headers).
		SetBody(body)

	var (
		resp *resty.Response
		err  error
	)

	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.Post(url)
	case "PUT":
		resp, err = req.Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		if responseStruct != nil {
			if err := json.Unmarshal(resp.Body(), responseStruct); err != nil {
				return fmt.Errorf("failed to unmarshal response body: %w", err)
			}
		}
		return nil
	}

	return fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode(), resp.Body())
}
