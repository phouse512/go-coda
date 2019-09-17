package coda

import (
	"fmt"
	"log"
)

type Row struct {
	Id          string                 `json:"id"`
	Type        string                 `json:"type"`
	Href        string                 `json:"href"`
	Name        string                 `json:"name"`
	Index       int                    `json:"index"`
	BrowserLink string                 `json:"browserLink"`
	CreatedAt   string                 `json:"createdAt"`
	UpdatedAt   string                 `json:"updatedAt"`
	Values      map[string]interface{} `json:"values"`
}

type ListRowsParameters struct {
	Query          string `json:"query"`
	SortBy         string `json:"sortBy"`
	UseColumnNames bool   `json:"useColumnNames"`
	PaginationPayload
}

type ListRowsResponse struct {
	Rows []Row `json:"items"`
	PaginationResponse
}

func (c *Client) ListTableRows(docId string, tableIdOrName string, listRowsParams ListRowsParameters) (ListRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, tableIdOrName)
	var rowsResp ListRowsResponse
	err := c.apiCall("GET", docPath, listRowsParams, &rowsResp)
	if err != nil {
		log.Print("Unable to list table rows with error.")
	}
	return rowsResp, err
}
