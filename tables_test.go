package coda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListTables(t *testing.T) {
	docId := "testDocId"
	expectedPath := fmt.Sprintf("/docs/%s/tables?tableTypes=table", docId)
	server := buildTestServer(expectedPath, "test_data/tables_list.json", 200, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	payload := ListTablesPayload{TableTypes: "table"}
	resp, err := testClient.ListTables(docId, payload)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp.Tables))
}
