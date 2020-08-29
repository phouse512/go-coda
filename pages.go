package coda

import (
	"fmt"
	"log"
)

type Icon struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	BrowserLink string `json:"browserLink"`
}

type Image struct {
	BrowserLink string `json:"browserLink"`
	Type        string `json:"type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type PageReference struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	BrowserLink string `json:"browserLink"`
	Href        string `json:"href"`
	Name        string `json:"name"`
}

type Page struct {
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	Href        string          `json:"href"`
	Name        string          `json:"name"`
	BrowserLink string          `json:"browserLink"`
	Children    []PageReference `json:"children"`
	Subtitle    string          `json:"subtitle"`
	Icon        Icon            `json:"icon"`
	Image       Image           `json:"image"`
	Parent      PageReference   `json:"parent"`
}

type ListPagesResponse struct {
	Pages []Page `json:"items"`
	PaginationResponse
}

type ListPagesPayload struct {
	Limit     int    `url:"limit,omitempty"`
	PageToken string `url:"pageToken,omitempty"`
}

type UpdatePagePayload struct {
	Name     string `json:"name"`
	Subtitle string `json:"subtitle"`
	IconName string `json:"iconName"`
	ImageUrl string `json:"imageUrl"`
}

type GetPageResponse struct {
	Page
}

type RequestResponse struct {
	RequestId string `json:"requestId"`
	Id        string `json:"id"`
}

func (c *Client) ListPages(docId string, paginationPayload PaginationPayload) (ListPagesResponse, error) {
	docPath := fmt.Sprintf("docs/%s/pages", docId)

	var pagesResp ListPagesResponse
	err := c.apiCall("GET", docPath, paginationPayload, &pagesResp)
	if err != nil {
		log.Print("Unable to list pages with error.")
	}

	return pagesResp, err
}

func (c *Client) GetPage(docId string, pageIdOrName string) (GetPageResponse, error) {
	docPath := fmt.Sprintf("docs/%s/pages/%s", docId, sectionIdOrName)

	var pageResponse GetPageResponse
	err := c.apiCall("GET", docPath, nil, &pageResponse)
	if err != nil {
		log.Print("Unable to get page with error.")
	}

	return pageResponse, err
}

func (c *Client) UpdatePage(docId string, pageIdOrName string, pagePayload UpdatePagePayload) (RequestResponse, error) {
	docPath := fmt.Sprintf("docs/%s/pages/%s", docId, pageIdOrName)
	var requestResp RequestResponse
	err := c.apiCall("PUT", docPath, pagePayload, &requestResp)
	if err != nil {
		log.Print("Unable to update page with error.")
	}
	return requestResp, err
}
