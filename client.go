// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/pkg/errors"
)

const (
	apiEndpoint = "https://codecov.io/api/"
	ghEndPoint  = "gh/"
)

// TestMethod is temp method
func (c *Client) TestMethod() {
	q := url.Values{
		"state": []string{"all"},
	}
	var res PullsResponse
	err := c.get("gh/budougumi0617/gopl/pulls", q, &res)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		panic(err)
	}
	fmt.Printf("a = %+v\n", res)
}

func newDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}

// HTTPClient is an interface for REST API.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var defaultHTTPClient HTTPClient = newDefaultHTTPClient()

// Client wraps http client
type Client struct {
	authToken string

	// HTTPClient is the HTTP client used for making requests against the
	// PagerDuty API. You can use either *http.Client here, or your own
	// implementation.
	HTTPClient HTTPClient
}

// NewClient creates an API client
func NewClient(authToken string) *Client {
	return &Client{
		authToken:  authToken,
		HTTPClient: defaultHTTPClient,
	}
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) put(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {

	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return c.do("PUT", path, bytes.NewBuffer(data), headers)
	}
	return c.do("PUT", path, nil, headers)
}

func (c *Client) post(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), headers)
}

func (c *Client) get(path string, query url.Values, v interface{}) error {
	if len(query) != 0 {
		path = path + "?" + query.Encode()
	}
	resp, err := c.do("GET", path, nil, nil)
	if err != nil {
		return err
	}
	return c.decodeJSON(resp, v)
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	// FIXME Need to able to change host site.
	endpoint := apiEndpoint + ghEndPoint + path
	req, _ := http.NewRequest(method, endpoint, body)
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+c.authToken)
	fmt.Println("get", endpoint)

	resp, err := c.HTTPClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(payload)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, errors.Wrap(err, "Error calling the API endpoint:")
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		return resp, errors.Errorf("Failed call API endpoint. HTTP response code: %v. response: %v", resp.StatusCode, resp)
	}
	return resp, nil
}
