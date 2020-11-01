package coda

import (
	"fmt"
	"log"
)

type DocCategory struct {
	Name string `json:"name"`
}

type GetDocumentCategoriesResponse struct {
	Categories []DocCategory `json:"items"`
}

type PublishDocumentPayload struct {
	Slug          string   `json:"slug"`
	Discoverable  bool     `json:"discoverable"`
	EarnCredit    bool     `json:"earnCredit"`
	CategoryNames []string `json:"categoryNames"`
	Mode          string   `json:"mode"`
}

type PublishDocumentResponse struct {
	RequestId string `json:"requestId"`
}

func (c *Client) GetDocumentCategories() (GetDocumentCategoriesResponse, error) {
	docPath := "/categories"

	var catResp GetDocumentCategoriesResponse
	err := c.apiCall("GET", docPath, nil, &catResp)
	if err != nil {
		log.Print("Unable to get document categories with error.")
	}

	return catResp, err
}

func (c *Client) PublishDocument(documentId string, payload PublishDocumentPayload) (PublishDocumentResponse, error) {
	docPath := fmt.Sprintf("/docs/%s/publish", documentId)

	var publishResp PublishDocumentResponse
	err := c.apiCall("PUT", docPath, payload, &publishResp)
	if err != nil {
		log.Print("Unable to publish document with error.")
	}

	return publishResp, err
}

func (c *Client) UnpublishDocument(documentId string) error {
	docPath := fmt.Sprintf("/docs/%s/publish", documentId)

	err := c.apiCall("DELETE", docPath, nil, nil)
	if err != nil {
		log.Print("Unable to unpublish document with error.")
	}

	return err
}
