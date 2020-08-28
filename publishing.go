package coda

import (
	"log"
)

type DocCategory struct {
	Name string `json:"name"`
}

type GetDocumentCategoriesResponse struct {
	Categories []DocCategory `json:"items"`
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
