package coda

import (
	"fmt"
	"log"
)

type Resource struct {
	Id   string `json:"id"`
	Href string `json:"href"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ResolveBrowserLinkParameters struct {
	Url               string `json:"url" url:"url"`
	DegradeGracefully bool   `json:"degradeGracefully" url:"degradeGracefully,omitempty"`
}

type ResolveBrowserLinkResponse struct {
	Type        string   `json:"type"`
	Href        string   `json:"href"`
	Resource    Resource `json:"resource"`
	BrowserLink string   `json:"browserLink"`
}

type MutationStatusResponse struct {
	IsCompleted bool `json:"completed"`
}

func (c *Client) ResolveBrowserLink(browserLinkParams ResolveBrowserLinkParameters) (ResolveBrowserLinkResponse, error) {
	docPath := fmt.Sprintf("resolveBrowserLink")
	var linkResp ResolveBrowserLinkResponse
	err := c.apiCall("GET", docPath, browserLinkParams, &linkResp)
	if err != nil {
		log.Print("Unable to resolve browser link.")
	}
	return linkResp, err
}

func (c *Client) GetMutationStatus(requestId string) (MutationStatusResponse, error) {
	docPath := fmt.Sprintf("mutationStatus/%s", requestId)
	var statusResp MutationStatusResponse
	err := c.apiCall("GET", docPath, nil, &statusResp)
	if err != nil {
		log.Print("Unable to get mutation status with error.")
	}

	return statusResp, err
}
