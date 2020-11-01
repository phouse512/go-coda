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
	Query          string `json:"query" url:"query"`
	SortBy         string `json:"sortBy" url:"sortBy"`
	UseColumnNames bool   `json:"useColumnNames" url:"useColumnNames"`
	ValueFormat    string `json:"valueFormat" url:"valueFormat"`
	VisibleOnly    bool   `json:"visibleOnly" url:"visibleOnly"`
	PaginationPayload
}

type GetRowParameters struct {
	UseColumnNames bool `json:"useColumnNames" url:"useColumnNames"`
}

type ListRowsResponse struct {
	Rows []Row `json:"items"`
	PaginationResponse
}

type GetRowResponse struct {
	Row
}

type CellParam struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

type RowParam struct {
	Cells []CellParam `json:"cells"`
}

type InsertRowsParameters struct {
	Rows       []RowParam `json:"rows"`
	KeyColumns []string   `json:"keyColumns,omitempty"`
}

type InsertRowsResponse struct {
	RequestId string `json:"requestId"`
}

type DeleteRowsParameters struct {
	RowIds []string `json:"rowIds"`
}

type DeleteRowsResponse struct {
	RowIds    []string `json:"rowIds"`
	RequestId string   `json:"requestId"`
}

type UpdateRowParameters struct {
	Row RowParam `json:"row"`
}

type RowQueryParams struct {
	DisableParsing bool `url:"disableParsing"`
}

type UpdateRowResponse struct {
	Id        string `json:"id"`
	RequestId string `json:"requestId"`
}

type DeleteRowResponse struct {
	Id        string `json:"id"`
	RequestId string `json:"requestId"`
}

type ListViewRowsParameters struct {
	Query          string `json:"query" url:"query"`
	SortBy         string `json:"sortBy" url:"sortBy"`
	UseColumnNames bool   `json:"useColumnNames" url:"useColumnNames"`
	ValueFormat    string `json:"valueFormat" url:"valueFormat"`
	PaginationPayload
}

type ListViewRowsResponse struct {
	Rows []Row `json:"items"`
	PaginationResponse
}

type PushButtonResponse struct {
	RequestId string `json:"requestId"`
	RowId     string `json:"rowId"`
	ColumnId  string `json:"columnId"`
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

func (c *Client) InsertRows(docId string, tableIdOrName string, disableParsing bool, insertRowParams InsertRowsParameters) (InsertRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, tableIdOrName)
	queryParams := RowQueryParams{DisableParsing: disableParsing}
	var rowsResp InsertRowsResponse
	err := c.apiCallFull("POST", docPath, insertRowParams, queryParams, &rowsResp)
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

func (c *Client) DeleteRows(docId string, tableIdOrName string, deleteRowsParams DeleteRowsParameters) (DeleteRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, tableIdOrName)
	var deleteResp DeleteRowsResponse
	err := c.apiCall("DELETE", docPath, deleteRowsParams, &deleteResp)
	if err != nil {
		log.Print("Unable to delete rows with error.")
	}
	return deleteResp, err
}

func (c *Client) UpdateRow(docId string, tableIdOrName string, rowIdOrName string, disableParsing bool, updateRowParams UpdateRowParameters) (UpdateRowResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows/%s", docId, tableIdOrName, rowIdOrName)
	queryParams := RowQueryParams{DisableParsing: disableParsing}
	var updateResp UpdateRowResponse
	err := c.apiCallFull("PUT", docPath, updateRowParams, queryParams, &updateResp)
	if err != nil {
		log.Print("Unable to update row with error.")
	}
	return updateResp, err
}

func (c *Client) DeleteRow(docId string, tableIdOrName string, rowIdOrName string) (DeleteRowResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows/%s", docId, tableIdOrName, rowIdOrName)
	var deleteResp DeleteRowResponse
	err := c.apiCall("DELETE", docPath, nil, &deleteResp)
	if err != nil {
		log.Print("Unable to delete row with error.")
	}
	return deleteResp, err
}

func (c *Client) ListViewRows(docId string, viewIdOrName string, viewRowsParams ListViewRowsParameters) (ListViewRowsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows", docId, viewIdOrName)
	var rowsResp ListViewRowsResponse
	err := c.apiCall("GET", docPath, viewRowsParams, &rowsResp)
	if err != nil {
		log.Print("Unable to get view rows with error.")
	}
	return rowsResp, err
}

func (c *Client) PushButton(docId string, tableIdOrName string, rowIdOrName string, columnIdOrName string) (PushButtonResponse, error) {
	docPath := fmt.Sprintf("docs/%s/tables/%s/rows/%s/buttons/%s", docId, tableIdOrName, rowIdOrName, columnIdOrName)
	var pushResp PushButtonResponse
	err := c.apiCall("POST", docPath, nil, &pushResp)
	if err != nil {
		log.Print("Unable to push button with error.")
	}
	return pushResp, err
}
