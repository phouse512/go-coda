package coda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

type GetDocumentResponse struct {
	Document Document
}

type ListDocumentsResponse struct {
	Documents []Document `json:"items"`
}

type Document struct {
	Id           string `json:"id"`
	DocumentType string `json:"type"`
	Href         string `json:"href"`
	BrowserLink  string `json:"browserLink"`
	Name         string `json:"name"`
	Owner        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func (c *Client) newRequest(method, methodPath string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: methodPath}
	rel.Path = path.Join(c.BaseURL.Path, rel.Path)
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return resp, buildError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) GetDoc(id string) (GetDocumentResponse, error) {
	docPath := fmt.Sprintf("/docs/%s", id)
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return GetDocumentResponse{}, err
	}

	var document Document
	resp, err := c.do(req, &document)
	if err != nil {
		log.Print("Unable to make request.")
		return GetDocumentResponse{}, err
	}
	log.Print("Received status: ", string(resp.StatusCode))
	var response = GetDocumentResponse{Document: document}
	return response, err
}

func (c *Client) ListDocs() ([]Document, error) {
	docPath := "/docs"
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return nil, err
	}

	var documentsResponse ListDocumentsResponse
	_, err = c.do(req, &documentsResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return nil, err
	}

	return documentsResponse.Documents, err
}
