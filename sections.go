package coda

import (
	"fmt"
	"log"
)

type Section struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Href        string `json:"href"`
	Name        string `json:"name"`
	BrowserLink string `json:"browserLink"`
	Parent      struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"parent"`
}

type ListSectionsResponse struct {
	Sections []Section `json:"items"`
	PaginationResponse
}

type ListSectionsPayload struct {
	Limit     int    `url:"limit,omitempty"`
	PageToken string `url:"pageToken,omitempty"`
}

type GetSectionResponse struct {
	Section
}

func (c *Client) ListSections(docId string, sectionsPayload ListSectionsPayload) (ListSectionsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/sections", docId)
	req, err := c.newRequest("GET", docPath, sectionsPayload)
	if err != nil {
		log.Print("Unable to create new request.")
		return ListSectionsResponse{}, err
	}

	var sectionsResponse ListSectionsResponse
	_, err = c.do(req, &sectionsResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return ListSectionsResponse{}, err
	}

	return sectionsResponse, err
}

func (c *Client) GetSection(docId string, sectionIdOrName string) (GetSectionResponse, error) {
	docPath := fmt.Sprintf("docs/%s/sections/%s", docId, sectionIdOrName)
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return GetSectionResponse{}, err
	}

	var sectionResponse GetSectionResponse
	_, err = c.do(req, &sectionResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return GetSectionResponse{}, err
	}

	return sectionResponse, err
}
