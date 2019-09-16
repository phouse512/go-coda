package coda

import (
	"fmt"
	"log"
)

type Control struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Href        string      `json:"href"`
	Name        string      `json:"name"`
	ControlType string      `json:"controlType"`
	Value       interface{} `json:"value"`
}

type ListControlsResponse struct {
	Controls []Control `json:"items"`
	PaginationResponse
}

type GetControlResponse struct {
	Control
}

func (c *Client) ListControls(docId string, paginationPayload PaginationPayload) (ListControlsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/controls", docId)
	var controlsResp ListControlsResponse
	err := c.apiCall("GET", docPath, paginationPayload, &controlsResp)
	if err != nil {
		log.Print("Unable to make api call with error.")
	}
	return controlsResp, err
}

func (c *Client) GetControl(docId string, controlIdOrName string) (GetControlResponse, error) {
	docPath := fmt.Sprintf("docs/%s/controls/%s", docId, controlIdOrName)
	var controlResp GetControlResponse
	err := c.apiCall("GET", docPath, nil, &controlResp)
	if err != nil {
		log.Print("Unable to make api call with error.")
	}
	return controlResp, err
}
