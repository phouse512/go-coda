package coda

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
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

func (c *Client) apiCall(method, url string, body interface{}, response interface{}) error {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return err
	}

	_, err = c.do(req, &response)
	if err != nil {
		log.Print("Unable to make request.")
		return err
	}
	return err
}

func (c *Client) apiCallFull(method, url string, body, queryParams, response interface{}) error {
	req, err := c.newRequestFull(method, url, body, queryParams)
	if err != nil {
		return err
	}

	_, err = c.do(req, &response)
	if err != nil {
		log.Print("Unable to make request.")
		return err
	}
	return err
}

func (c *Client) newRequestFull(method, methodPath string, body interface{}, queryParams interface{}) (*http.Request, error) {
	rel := &url.URL{Path: methodPath}
	rel.Path = path.Join(c.BaseURL.Path, rel.Path)
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	encodedParams, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	u.RawQuery = encodedParams.Encode()

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

func (c *Client) newRequest(method, methodPath string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: methodPath}
	rel.Path = path.Join(c.BaseURL.Path, rel.Path)
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil && (method == "POST" || method == "DELETE" || method == "PUT") {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	if body != nil && method == "GET" {
		queryParams, _ := query.Values(body)
		u.RawQuery = queryParams.Encode()
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
