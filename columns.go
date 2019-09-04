package coda

import (
	"fmt"
	"log"
)

type Column struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Href         string `json:"href"`
	Name         string `json:"name"`
	IsDisplay    bool   `json:"display"`
	IsCalculated bool   `json:"calculated"`
}

type ListColumnsResponse struct {
	Columns []Column `json:"items"`
	PaginationResponse
}

type GetColumnResponse struct {
	Column
}

func (c *Client) ListColumns(docId string, tableIdOrName string, paginationPayload PaginationPayload) (ListColumnsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/columns", docId, tableIdOrName)
	var columnsResponse ListColumnsResponse
	err := c.apiCall("GET", docPath, paginationPayload, &columnsResponse)
	if err != nil {
		log.Print("Unable to get columns with error.")
		return columnsResponse, err
	}
	return columnsResponse, err
}

func (c *Client) GetColumn(docId string, tableIdOrName string, columnIdOrName string) (GetColumnResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/columns/%s", docId, tableIdOrName, columnIdOrName)
	var columnResponse GetColumnResponse
	err := c.apiCall("GET", docPath, nil, &columnResponse)
	if err != nil {
		log.Print("Unable to get column with error.")
		return columnResponse, err
	}
	return columnResponse, err
}

func (c *Client) ListViewColumns(docId string, viewIdOrName string, paginationPayload PaginationPayload) (ListColumnsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/views/%s/columns", docId, viewIdOrName)
	var columnsResp ListColumnsResponse
	err := c.apiCall("GET", docPath, paginationPayload, &columnsResp)
	if err != nil {
		log.Print("Unable to list view columns with error.")
		return columnsResp, err
	}
	return columnsResp, err
}
