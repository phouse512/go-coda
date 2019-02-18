package coda

type PaginationResponse struct {
	Href          string `json:"href,omitempty"`
	NextPageToken string `json:"nextPageToken,omitempty"`
	NextPageLink  string `json:"nextPageLink,omitempty"`
}
