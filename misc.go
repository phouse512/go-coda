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
	Url               string `json:"url"`
	DegradeGracefully bool   `json:"degradeGracefully"`
}

type ResolveBrowserLinkResponse struct {
	Type        string   `json:"type"`
	Href        string   `json:"href"`
	Resource    Resource `json:"resource"`
	BrowserLink string   `json:"browserLink"`
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
