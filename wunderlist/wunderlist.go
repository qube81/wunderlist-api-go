package wunderlist

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	endpoint = "https://a.wunderlist.com/api/v1/"
)

var httpClient = &http.Client{
	Timeout: time.Duration(5) * time.Second,
}

// Client for wunderlist API
type Client struct {
	clientID    string
	accessToken string
	httpClient  *http.Client
}

// NewClient generate and return Client
func NewClient(clientID string, accessToken string) *Client {

	c := &Client{
		clientID:    clientID,
		accessToken: accessToken,
		httpClient:  httpClient,
	}

	return c
}

// Get request HTTP GET via Client
func (c *Client) Get(path string) (body string, err error) {

	req, err := http.NewRequest("GET", endpoint+path, nil)
	if err != nil {
		return "", err
	}
	return Execute(c, req)
}

// Execute HTTP request
func Execute(c *Client, req *http.Request) (body string, err error) {

	req.Header.Add("X-Client-ID", c.clientID)
	req.Header.Add("X-Access-Token", c.accessToken)

	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), err
}
