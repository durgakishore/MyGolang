package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client  *http.Client
	Headers http.Header

	maxIdleConnections int
	connectionTimeout  time.Duration
	requestTimeout     time.Duration
}

//NewHTTPClient returns HTTP inteface object
func NewHTTPClient() HTTPClient {
	httpClient := &httpClient{}

	return httpClient
}

//HTTPClient interface
type HTTPClient interface {
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetRequestTimeout(timeout time.Duration)
	SetMaxIdleConnections(i int)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

//SetHeaders to set the headers for the request
func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}
func (c *httpClient) SetRequestTimeout(timeout time.Duration) {
	c.requestTimeout = timeout
}
func (c *httpClient) SetMaxIdleConnections(i int) {
	c.maxIdleConnections = i
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {

	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {

	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {

	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {

	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
