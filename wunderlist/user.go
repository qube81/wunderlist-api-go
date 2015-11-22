package wunderlist

import (
	"net/url"
)

//User for WunderList
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Revision  int    `json:"revision"`
	CreatedAt string `json:"created_at"`
}

// UserAPI https://developer.wunderlist.com/documentation/endpoints/user
type UserAPI struct {
	client *Client
}

// Get fetch the currently logged in user
func (a *UserAPI) Get() (result User, err error) {
	var user User
	if err := a.client.Get("user", &user, url.Values{}); err != nil {
		return user, err
	}

	return user, nil
}
