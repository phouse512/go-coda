package coda

type Row struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Href        string `json:"href"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	BrowserLink string `json:"browserLink"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
