package coda

import (
	"fmt"
	"log"
)

type Folder struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Href     string    `json:"href"`
	Name     string    `json:"name"`
	Children []Section `json:"children"`
}

type ListFoldersResponse struct {
	Folders []Folder `json:"items"`
	PaginationResponse
}

type GetFolderResponse struct {
	Folder
}

func (c *Client) ListFolders(docId string, paginationPayload PaginationPayload) (ListFoldersResponse, error) {
	docPath := fmt.Sprintf("docs/%s/folders", docId)
	req, err := c.newRequest("GET", docPath, paginationPayload)
	if err != nil {
		log.Print("Unable to create new request.")
		return ListFoldersResponse{}, err
	}

	var foldersResponse ListFoldersResponse
	_, err = c.do(req, &foldersResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return ListFoldersResponse{}, err
	}

	return foldersResponse, err
}

func (c *Client) GetFolder(docId string, folderIdOrName string) (GetFolderResponse, error) {
	docPath := fmt.Sprintf("docs/%s/folders/%s", docId, folderIdOrName)
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return GetFolderResponse{}, err
	}

	var folderResponse GetFolderResponse
	_, err = c.do(req, &folderResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return GetFolderResponse{}, err
	}

	return folderResponse, err
}
