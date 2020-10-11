package coda

import (
	"fmt"
	"log"
)

type Principal struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Permission struct {
	Principal Principal `json:"principal"`
	Id        string    `json:"id"`
	Access    string    `json:"access"`
}

type AddPermissionPayload struct {
	Access        string    `json:"access"`
	Principal     Principal `json:"principal"`
	SuppressEmail bool      `json:"suppressEmail"`
}

type GetACLResponse struct {
	CanShare        bool `json:"canShare"`
	CanShareWithOrg bool `json:"canShareWithOrg"`
}

type ListACLResponse struct {
	Permissions []Permission `json:"items"`
	PaginationResponse
}

type AddPermissionResponse struct{}

type DeletePermissionResponse struct{}

func (c *Client) GetACLMetadata(docId string) (GetACLResponse, error) {
	docPath := fmt.Sprintf("/docs/%s/acl/metadata", docId)

	var aclResp GetACLResponse
	err := c.apiCall("GET", docPath, nil, &aclResp)
	if err != nil {
		log.Print("Unable to get document acl with error.")
	}

	return aclResp, err
}

func (c *Client) ListACLs(docId string, paginationPayload PaginationPayload) (ListACLResponse, error) {
	docPath := fmt.Sprintf("/docs/%s/acl/permissions", docId)

	var resp ListACLResponse
	err := c.apiCall("GET", docPath, paginationPayload, &resp)
	if err != nil {
		log.Print("Unable to list permissions with error.")
	}

	return resp, err
}

func (c *Client) AddPermission(docId string, payload AddPermissionPayload) (AddPermissionResponse, error) {
	docPath := fmt.Sprintf("/docs/%s/acl/permissions", docId)

	var resp AddPermissionResponse
	err := c.apiCall("POST", docPath, payload, &resp)
	if err != nil {
		log.Print("Unable to add permission with error.")
	}

	return resp, err
}

func (c *Client) DeletePermission(docId, permissionId string) (DeletePermissionResponse, error) {
	docPath := fmt.Sprintf("/docs/%s/acl/permissions/%s", docId, permissionId)

	var resp DeletePermissionResponse
	err := c.apiCall("DELETE", docPath, nil, &resp)
	if err != nil {
		log.Print("Unable to delete permission with error.")
	}
	return resp, err
}
