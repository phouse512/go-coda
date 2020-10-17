package coda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDoc(t *testing.T) {
	expectedPath := fmt.Sprintf("/docs")
	server := buildTestServer(expectedPath, "test_data/docs_create.json", 201, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	docPayload := CreateDocumentPayload{
		Title:     "fakeTitle",
		SourceDoc: "dskfd",
		Timezone:  "America/Chicago",
		FolderId:  "kdfjlaksdf",
	}
	_, err := testClient.CreateDocument(docPayload)
	assert.Nil(t, err)
}

func TestListDocs(t *testing.T) {
	expectedPath := fmt.Sprintf("/docs")
	server := buildTestServer(expectedPath, "test_data/list_docs.json", 200, t)
	defer server.Close()
	testClient := buildTestClient(server.URL)

	docsResp, err := testClient.ListDocuments(ListDocumentsPayload{})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(docsResp.Documents))
	assert.Equal(t, "AbCDeFGH", docsResp.Documents[0].Id)

	assert.Equal(t, "20", docsResp.PaginationResponse.NextPageToken)
}
