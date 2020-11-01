package coda

import (
	"fmt"
	"log"
)

type ControlReference struct {
	Id     string        `json:"id"`
	Type   string        `json:"type"`
	Href   string        `json:"href"`
	Name   string        `json:"name"`
	Parent PageReference `json:"parent"`
}

type Control struct {
	ControlReference
	ControlType string      `json:"controlType"`
	Value       interface{} `json:"value"`
}

type ListControlsResponse struct {
	Controls []ControlReference `json:"items"`
	PaginationResponse
}

type ListControlsPayload struct {
	SortBy string `json:"sortBy" url:"sortBy,omitEmpty"`
	PaginationPayload
}

type GetControlResponse struct {
	Control
}

func (c *Client) ListControls(docId string, payload ListControlsPayload) (ListControlsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/controls", docId)
	var controlsResp ListControlsResponse
	err := c.apiCall("GET", docPath, payload, &controlsResp)
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
