package wunderlist

import (
	"net/url"
	"strconv"
)

// https://developer.wunderlist.com/documentation/endpoints/list

//List for WunderList
type List struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	OwnerType          string `json:"owner_type"`
	OwnerID            int    `json:"owner_id"`
	ListType           string `json:"list_type"`
	Public             bool   `json:"public"`
	Revision           int    `json:"revision"`
	CreatedAt          string `json:"created_at"`
	CreatedByRequestID string `json:"created_by_request_id"`
	Type               string `json:"type"`
}

//Lists for WunderList
type Lists struct {
	Collection []List
}

// ListAPI https://developer.wunderlist.com/documentation/endpoints/list
type ListAPI struct {
	client *Client
}

// GetAll get all Lists a user has permission to
func (a *ListAPI) GetAll() (result []List, err error) {
	var lists []List
	if err := a.client.Get("lists", &lists, url.Values{}); err != nil {
		return lists, err
	}

	return lists, nil
}

// Get get a specific List
func (a *ListAPI) Get(id int) (result List, err error) {
	var list List
	if err := a.client.Get("lists/"+strconv.Itoa(id), &list, url.Values{}); err != nil {
		return list, err
	}

	return list, nil
}
