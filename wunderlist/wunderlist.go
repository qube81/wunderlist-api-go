package wunderlist

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// https://developer.wunderlist.com/documentation

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
	List        *ListAPI
	User        *UserAPI
	Task        *TaskAPI
}

// NewClient generate and return Client
func NewClient(clientID string, accessToken string) *Client {

	c := &Client{}
	c.clientID = clientID
	c.accessToken = accessToken
	c.httpClient = httpClient
	c.List = &ListAPI{client: c}
	c.User = &UserAPI{client: c}
	c.Task = &TaskAPI{client: c}

	return c
}

// Get request HTTP GET via Client
func (c *Client) Get(path string, v interface{}, query url.Values) (err error) {

	req, err := http.NewRequest("GET", endpoint+path, nil)
	if err != nil {
		return err
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return Execute(c, req, v)
}

// Post request HTTP POST via Client
func (c *Client) Post(path string, v interface{}, json string) (err error) {
	req, err := http.NewRequest("POST", endpoint+path, strings.NewReader(json))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return err
	}

	return Execute(c, req, v)
}

// Patch request HTTP POST via Client
func (c *Client) Patch(path string, v interface{}, json string) (err error) {
	req, err := http.NewRequest("PATCH", endpoint+path, strings.NewReader(json))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return err
	}

	return Execute(c, req, v)
}

// Execute HTTP request
func Execute(c *Client, req *http.Request, v interface{}) (err error) {

	req.Header.Add("X-Client-ID", c.clientID)
	req.Header.Add("X-Access-Token", c.accessToken)
	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New(string(respBody))
	}

	return json.Unmarshal(respBody, v)
}
