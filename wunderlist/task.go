package wunderlist

import (
	_ "github.com/k0kubun/pp"
	"net/url"
	"strconv"
)

//Task for WunderList
type Task struct {
	ID                 int    `json:"id"`
	CreatedAt          string `json:"created_at"`
	CreatedByID        int    `json:"created_by_id"`
	CreatedByRequestID string `json:"created_by_request_id"`
	DueDate            string `json:"due_date"`
	ListID             int    `json:"list_id"`
	Starred            bool   `json:"starred"`
	Completed          bool   `json:"completed"`
	Title              string `json:"title"`
	Revision           int    `json:"revision"`
	Type               string `json:"type"`
}

// Tasks for WunderList
type Tasks struct {
	Collection []Task
}

// TaskAPI https://developer.wunderlist.com/documentation/endpoints/task
type TaskAPI struct {
	client *Client
}

// GetByListID fetch the currently logged in user
func (a *TaskAPI) GetByListID(listID int, completed ...bool) (result []Task, err error) {
	var tasks []Task
	values := url.Values{}
	values.Add("list_id", strconv.Itoa(listID))

	if len(completed) > 0 {
		values.Add("completed", strconv.FormatBool(completed[0]))
	}

	if err := a.client.Get("tasks", &tasks, values); err != nil {
		return tasks, err
	}
	return tasks, nil
}
