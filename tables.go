package coda

import (
	"fmt"
	"log"
)

type TableFilter struct {
	Valid           bool `json:"valid"`
	IsVolatile      bool `json:"isVolatile"`
	HasUserFormula  bool `json:"hasUserFormula"`
	HasTodayFormula bool `json:"hasTodayFormula"`
	HasNowFormula   bool `json:"hasNowFormula"`
}

type Table struct {
	TableReference
	DisplayColumn Column `json:"displayColumn"`
	RowCount      int    `json:"rowCount"`
	Sorts         []struct {
		Column    Column `json:"column"`
		Direction string `json:"direction"`
	} `json:"sorts"`
	Layout      string         `json:"layout"`
	CreatedAt   string         `json:"createdAt"`
	UpdatedAt   string         `json:"updatedAt"`
	ParentTable TableReference `json:"parentTable"`
	Filter      TableFilter    `json:"filter"`
}

type TableReference struct {
	Id          string        `json:"id"`
	Type        string        `json:"type"`
	TableType   string        `json:"tableType"`
	BrowserLink string        `json:"browserLink"`
	Href        string        `json:"href"`
	Name        string        `json:"name"`
	Parent      PageReference `json:"parent"`
}

type ListTablesResponse struct {
	Tables []TableReference `json:"items"`
	PaginationResponse
}

type ListTablesPayload struct {
	SortBy     string `url:"sortBy,omitempty"`
	TableTypes string `url:"tableTypes,omitempty"`
	PaginationPayload
}

type GetTableResponse struct {
	Table
}

func (c *Client) ListTables(docId string, payload ListTablesPayload) (ListTablesResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables", docId)

	var tablesResp ListTablesResponse
	err := c.apiCall("GET", docPath, payload, &tablesResp)
	if err != nil {
		log.Print("Unable to list tables with error.")
	}

	return tablesResp, err
}

func (c *Client) GetTable(docId string, tableIdOrName string) (GetTableResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s", docId, tableIdOrName)

	var tableResp GetTableResponse
	err := c.apiCall("GET", docPath, nil, &tableResp)
	if err != nil {
		log.Print("Unable to get table with error.")
	}

	return tableResp, err
}
