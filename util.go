package coda

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func buildTestClient(testServerUrl string) *Client {
	httpClient := &http.Client{
		Transport: &transport{
			defaultTransport: http.DefaultTransport,
			token:            "fakeKey",
		},
	}

	u, _ := url.Parse(testServerUrl)
	codaClient := &Client{
		UserAgent:  "golang_bot/1.0",
		HttpClient: httpClient,
		BaseURL:    u,
	}

	return codaClient
}

func buildTestServer(expectedPath, sampleDataPath string, expectedStatus int, t *testing.T) *httptest.Server {
	// load data from filepath
	data, err := ioutil.ReadFile(sampleDataPath)
	if err != nil {
		panic(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), expectedPath)
		rw.WriteHeader(expectedStatus)
		rw.Write(data)
	}))
	return server
}
