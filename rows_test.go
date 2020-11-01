package coda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListRows(t *testing.T) {
	docId := "abc"
	tableId := "123"
	expectedPath := fmt.Sprintf("/docs/%s/tables/%s/rows?query=colId%%3A131231&sortBy=&useColumnNames=false&valueFormat=&visibleOnly=true", docId, tableId)
	server := buildTestServer(expectedPath, "test_data/rows_list.json", 200, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	params := ListRowsParameters{
		Query:       "colId:131231",
		VisibleOnly: true,
	}
	resp, err := testClient.ListTableRows(docId, tableId, params)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp.Rows))
	assert.Equal(t, 7, resp.Rows[0].Index)
}

func TestInsertRows(t *testing.T) {
	docId := "abc"
	tableId := "123"
	expectedPath := fmt.Sprintf("/docs/%s/tables/%s/rows?disableParsing=true", docId, tableId)
	expectedBody := "{\"rows\":[{\"cells\":[{\"column\":\"DateColumn\",\"value\":\"10/12/23\"},{\"column\":\"MoneyColumn\",\"value\":1.23}]}],\"keyColumns\":[\"dateColumnId\"]}\n"
	server := buildTestServerFull(expectedPath, "test_data/rows_insert.json", 202, expectedBody, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	cellParams := []CellParam{
		CellParam{
			Column: "DateColumn",
			Value:  "10/12/23",
		},
		CellParam{
			Column: "MoneyColumn",
			Value:  1.23,
		},
	}
	params := InsertRowsParameters{
		Rows: []RowParam{
			RowParam{
				Cells: cellParams},
		},
		KeyColumns: []string{"dateColumnId"},
	}

	_, err := testClient.InsertRows(docId, tableId, true, params)
	assert.Nil(t, err)
}
