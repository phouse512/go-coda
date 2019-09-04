package coda

import (
	"fmt"
	"log"
)

type Table struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	Href          string `json:"href"`
	Name          string `json:"name"`
	BrowserLink   string `json:"browserLink"`
	DisplayColumn Column `json:"displayColumn"`
	RowCount      int    `json:"rowCount"`
	Sorts         []struct {
		Column    Column `json:"column"`
		Direction string `json:"direction"`
	} `json:"sorts"`
	Layout    string `json:"layout"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ListTablesResponse struct {
	Tables []Table `json:"items"`
	PaginationResponse
}

type GetTableResponse struct {
	Table
}

func (c *Client) ListTables(docId string, paginationPayload PaginationPayload) (ListTablesResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables", docId)
	req, err := c.newRequest("GET", docPath, paginationPayload)
	if err != nil {
		log.Print("Unable to create new request.")
		return ListTablesResponse{}, err
	}

	var tablesResponse ListTablesResponse
	_, err = c.do(req, &tablesResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return ListTablesResponse{}, err
	}

	return tablesResponse, err
}

func (c *Client) GetTable(docId string, tableIdOrName string) (GetTableResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s", docId, tableIdOrName)
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return GetTableResponse{}, err
	}

	var tableResponse GetTableResponse
	_, err = c.do(req, &tableResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return GetTableResponse{}, err
	}

	return tableResponse, err
}
