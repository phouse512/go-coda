package coda

type User struct {
	Name    string `json:"name"`
	LoginId string `json:"loginId"`
	Type    string `json:"type"`
	Href    string `json:"href"`
}
