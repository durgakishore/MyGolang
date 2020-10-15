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

	if finaheaders.Get("Content-Type") != "application/json" {
		t.Error("Received an invalid content type")
	}

	if finaheaders.Get("X-Request-ID") != "ABC-123" {
		t.Error("Received an invalid X-Request-ID")
	}
}

func TestGetRequestBody(t *testing.T) {

	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {

		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body  expected when passing nil body")
		}

	})

	t.Run("BodyWithJson", func(t *testing.T) {

		//Execution

		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		//Validation

		if err != nil {
			t.Error("no error expected when marshalling a slice")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body observed")
		}

	})

	t.Run("BodyWithXML", func(t *testing.T) {

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

	})
}
