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
	Parent      struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Href string `json"href"`
	} `json:"parent"`
}

type ListRowsParameters struct {
	Query          string `json:"query"`
	SortBy         string `json:"sortBy"`
	UseColumnNames bool   `json:"useColumnNames"`
	PaginationPayload
}

type GetRowParameters struct {
	UseColumnNames bool `json:"useColumnNames"`
}

type ListRowsResponse struct {
	Rows []Row `json:"items"`
	PaginationResponse
}

type GetRowResponse struct {
	Row
}

type InsertRowsParameters struct {
	Rows []struct {
		Cells []struct {
			Column string      `json:"column"`
			Value  interface{} `json:"value"`
		} `json:"cells"`
	} `json:"rows"`
}

type InsertRowsResponse struct{}

func (c *Client) ListTableRows(docId string, tableIdOrName string, listRowsParams ListRowsParameters) (ListRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, tableIdOrName)
	var rowsResp ListRowsResponse
	err := c.apiCall("GET", docPath, listRowsParams, &rowsResp)
	if err != nil {
		log.Print("Unable to list table rows with error.")
	}
	return rowsResp, err
}

func (c *Client) InsertRows(docId string, tableIdOrName string, insertRowParams InsertRowsParameters) (InsertRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, tableIdOrName)
	var rowsResp InsertRowsResponse
	err := c.apiCall("POST", docPath, insertRowParams, &rowsResp)
	if err != nil {
		log.Print("Unable to insert rows with error.")
	}
	return rowsResp, err
}

func (c *Client) GetTableRow(docId string, tableIdOrName string, rowIdOrName string, getRowParams GetRowParameters) (GetRowResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows/%s", docId, tableIdOrName, rowIdOrName)
	var rowResp GetRowResponse
	err := c.apiCall("GET", docPath, getRowParams, &rowResp)
	if err != nil {
		log.Print("Unable to get table row with error.")
	}
	return rowResp, err
}
