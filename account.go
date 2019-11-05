package coda

import (
	"fmt"
	"log"
)

type UserInfoResponse struct {
	Name      string `json:"name"`
	LoginId   string `json:"loginId"`
	Type      string `json:"type"`
	Scoped    bool   `json:"scoped"`
	TokenName string `json:"tokenName"`
	Href      string `json:"href"`
}

func (c *Client) GetUserInfo() (UserInfoResponse, error) {
	docPath := fmt.Sprintf("whoami")
	var userResp UserInfoResponse
	err := c.apiCall("GET", docPath, nil, &userResp)
	if err != nil {
		log.Print("Unable to get user info with error.")
	}
	return userResp, err
}
