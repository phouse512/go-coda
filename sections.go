package coda

import (
//	"log"
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
}

type ListSectionsPayload struct {
	Limit     int    `schema:"limit"`
	PageToken string `schema:"pageToken"`
}
