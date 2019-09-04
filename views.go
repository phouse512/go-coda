package coda

import (
	"fmt"
	"log"
)

type View struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	Href          string `json:"href"`
	Name          string `json:"name"`
	BrowserLink   string `json:"browserLink"`
	DisplayColumn Column `json:"displayColumn"`
	RowCount      int    `json:"rowCount"`
	Sorts         []struct {
		Column    Column `json:"column"`
		Direction string `json:"direction"`
	}
	Layout    string `json:"layout"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ListViewsResponse struct {
	Views []View `json:"items"`
	PaginationResponse
}

type GetViewResponse struct {
	View
}

func (c *Client) ListViews(docId string, paginationPayload PaginationPayload) (ListViewsResponse, error) {
	docPath := fmt.Sprintf("docs/%s/views", docId)
	var viewsResponse ListViewsResponse
	err := c.apiCall("GET", docPath, paginationPayload, &viewsResponse)
	if err != nil {
		log.Print("Unable to make api call with error.")
		return viewsResponse, err
	}
	return viewsResponse, err
}

func (c *Client) GetView(docId string, viewIdOrName string) (GetViewResponse, error) {
	docPath := fmt.Sprintf("docs/%s/views/%s", docId, viewIdOrName)
	req, err := c.newRequest("GET", docPath, nil)
	if err != nil {
		log.Print("Unable to create new request.")
		return GetViewResponse{}, err
	}

	var viewResponse GetViewResponse
	_, err = c.do(req, &viewResponse)
	if err != nil {
		log.Print("Unable to make request.")
		return GetViewResponse{}, err
	}

	return viewResponse, err
}
