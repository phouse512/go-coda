package coda

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func BuildTestClient(testServerUrl string) *Client {
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

func TestDeletePermission(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/docs/fakeDoc/acl/permissions/fakePermission")
		rw.Write([]byte(`{}`))
	}))
	defer server.Close()

	testClient := BuildTestClient(server.URL)

	resp, err := testClient.DeletePermission("fakeDoc", "fakePermission")
	log.Print(err)
	log.Print(resp)
}
