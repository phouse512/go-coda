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
}

type ListSectionsResponse struct {
	Sections []Section `json:"items"`
	PaginationResponse
}

type ListSectionsPayload struct {
	Limit     int    `schema:"limit"`
	PageToken string `schema:"pageToken"`
}

func (c *Client) ListSections(docId string) (ListSectionsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/sections", docId)
	req, err := c.newRequest("GET", docPath, nil)
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
