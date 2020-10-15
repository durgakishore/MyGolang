package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {

	//Initialization

	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")

	client.Headers = commonHeaders

	//Execution

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-ID", "ABC-123")
	finaheaders := client.getRequestHeaders(requestHeaders)

	//Validation

	if len(finaheaders) != 3 {
		t.Error("We expect 3 headers")
	}
}
