package coda

import (
	"log"
)

type DocPublished struct {
	BrowserLink  string        `json:"browserLink"`
	Discoverable bool          `json:"discoverable"`
	EarnCredit   bool          `json:"earnCredit"`
	Mode         string        `json:"mode"`
	Categories   []DocCategory `json:"categories"`
	Description  string        `json:"description"`
	ImageLink    string        `json:"imageLink"`
}

type DocSize struct {
	TotalRowCount     int  `json:"totalRowCount"`
	TableAndViewCount int  `json:"tableAndViewCount"`
	PageCount         int  `json:"pageCount"`
	OverApiSizeLimit  bool `json:"overApiSizeLimit"`
}

type Document struct {
	Id          string  `json:"id"`
	Type        string  `json:"type"`
	Href        string  `json:"href"`
	BrowserLink string  `json:"browserLink"`
	Name        string  `json:"name"`
	Owner       string  `json:"email"`
	OwnerName   string  `json:"ownerName"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	Icon        Icon    `json:"icon"`
	DocSize     DocSize `json:"docSize"`
	SourceDoc   struct {
		Id          string `json:"id"`
		Type        string `json:"type"`
		BrowserLink string `json:"browserLink"`
		Href        string `json:"href"`
	}
	Published DocPublished `json:"published"`
}

type GetDocumentResponse struct {
	Document
}

type DeleteDocumentResponse struct{}

type ListDocumentsPayload struct {
	IsOwner        bool   `url:"isOwner,omitempty"`
	Query          string `url:"query,omitempty"`
	SourceDocument string `url:"sourceDoc,omitempty"`
	IsStarred      bool   `url:"isStarred,omitempty"`
	InGallery      bool   `url:"inGallery,omitempty"`
	WorkspaceId    string `url:"workspaceId,omitempty"`
	FolderId       string `url:"folderId,omitempty"`
	PaginationPayload
}

type ListDocumentsResponse struct {
	Documents []Document `json:"items"`
	PaginationResponse
}

type CreateDocumentPayload struct {
	Title     string `json:"title,omitempty"`
	SourceDoc string `json:"sourceDoc,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	FolderId  string `json:"folderId,omitempty"`
}

type CreateDocumentResponse struct {
	Document
}

func (c *Client) GetDoc(id string) (GetDocumentResponse, error) {
	docPath := fmt.Sprintf("/docs/%s", id)

	var docResp GetDocumentResponse
	err := c.apiCall("GET", docPath, nil, &docResp)
	if err != nil {
		log.Print("Unable to get document with error.")
	}

	return docResp, err
}

func (c *Client) ListDocuments(payload ListDocumentsPayload) (ListDocumentsResponse, error) {
	docPath := "/docs"

	var docsResp ListDocumentsResponse
	err := c.apiCall("GET", docPath, payload, &docsResp)
	if err != nil {
		log.Print("Unable to list documents with error.")
	}

	return docsResp, err
}

func (c *Client) CreateDocument(payload CreateDocumentPayload) (CreateDocumentResponse, error) {
	docPath := "/docs"
	var docResp CreateDocumentResponse
	err := c.apiCall("POST", docPath, payload, &docResp)
	if err != nil {
		log.Print("Unable to create document with error.")
	}

	return docResp, err
}

func (c *Client) DeleteDocument(id string) (DeleteDocumentResponse, error) {
	docPath := fmt.Sprintf("/docs/%s", id)

	var docResp DeleteDocumentResponse
	err := c.apiCall("DELETE", docPath, nil, &docResp)
	if err != nil {
		log.Print("Unable to delete document with error.")
	}

	return docResp, err
}
