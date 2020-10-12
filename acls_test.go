package coda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

func BuildTestServer(expectedPath, sampleDataPath string, expectedStatus int, t *testing.T) *httptest.Server {
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

func TestDeletePermission(t *testing.T) {
	docId := "fakeDoc"
	permission := "fakePermission"
	expectedPath := fmt.Sprintf("/docs/%s/acl/permissions/%s", docId, permission)
	server := BuildTestServer(expectedPath, "test_data/delete_permission.json", 200, t)
	defer server.Close()
	testClient := BuildTestClient(server.URL)

	_, err := testClient.DeletePermission(docId, permission)
	assert.Nil(t, err)
}
