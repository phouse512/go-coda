package coda

type PaginationResponse struct {
	Href          string `json:"href,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	NextPageLink  string `json:"nextPageLink,omitempty"`
}

type PaginationPayload struct {
	Limit     int    `url:"limit,omitempty"`
	PageToken string `url:"pageToken,omitempty"`
}
