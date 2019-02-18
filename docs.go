package coda

import (
	"log"
)

type CreateDocPayload struct {
	Title     string `json:"title"`
	SourceDoc string `json:"sourceDoc,omitempty"`
}

type CreateDocResponse struct {
	Document Document
}

func (c *Client) CreateDoc(payload CreateDocPayload) (CreateDocResponse, error) {
	docPath := "/docs"
	req, err := c.newRequest("POST", docPath, payload)
	if err != nil {
		log.Print("Unable to create new request.")
		return CreateDocResponse{}, err
	}

	var document Document
	_, err = c.do(req, &document)
	if err != nil {
		log.Print("Unable to make request.")
		return CreateDocResponse{}, err
	}

	return CreateDocResponse{Document: document}, nil
}
