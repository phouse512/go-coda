package coda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeletePermission(t *testing.T) {
	docId := "fakeDoc"
	permission := "fakePermission"
	expectedPath := fmt.Sprintf("/docs/%s/acl/permissions/%s", docId, permission)
	server := buildTestServer(expectedPath, "test_data/delete_permission.json", 200, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	_, err := testClient.DeletePermission(docId, permission)
	assert.Nil(t, err)
}
