package coda

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

type Document struct {
	id           string `json:"id"`
	documentType string `json:"type"`
	href         string `json:"href"`
	browserLink  string `json:"href"`
	name         string `json:"href"`
	owner        string `json:"email"`
	createdAt    string `json:"createdAt"`
	updatedAt    string `json:"updatedAt"`
}

func (c *Client) ListDocs() ([]Document, error) {
	rel := &url.URL{Path: "/docs"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var docs []Document
	err = json.NewDecoder(resp.Body).Decode(&docs)
	return docs, err
}
