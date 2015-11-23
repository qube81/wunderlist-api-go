package wunderlist

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
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

// CreateTask for POST
type CreateTask struct {
	ListID          int    `json:"list_id"`
	Title           string `json:"title"`
	AssigneeID      int    `json:"assignee_id"`
	Completed       bool   `json:"completed"`
	Starred         bool   `json:"starred"`
	DueDate         string `json:"due_date"`
	RecurrenceType  string `json:"recurrence_type"`  //must be accompanied by recurrence_count
	RecurrenceCount int    `json:"recurrence_count"` //must be accompanied by recurrence_type
}

// UpdateTask for Patch
type UpdateTask struct {
	Revision        int      `json:"revision"`
	Title           string   `json:"title"`
	AssigneeID      int      `json:"assignee_id"`
	Completed       bool     `json:"completed"`
	Starred         bool     `json:"starred"`
	DueDate         string   `json:"due_date"`
	RecurrenceType  string   `json:"recurrence_type"`  //must be accompanied by recurrence_count
	RecurrenceCount int      `json:"recurrence_count"` //must be accompanied by recurrence_type
	Remove          []string `json:"remove"`
}

// TaskAPI https://developer.wunderlist.com/documentation/endpoints/task
type TaskAPI struct {
	client *Client
}

// GetByListID Tasks for a List with completed option
func (a *TaskAPI) GetByListID(listID int, completed bool) (result []Task, err error) {
	var tasks []Task
	query := url.Values{}
	query.Add("list_id", strconv.Itoa(listID))
	query.Add("completed", strconv.FormatBool(completed))

	if err := a.client.Get("tasks", &tasks, query); err != nil {
		return tasks, err
	}
	return tasks, nil
}

// Get get a specific task
func (a *TaskAPI) Get(id int) (result Task, err error) {
	var task Task
	if err := a.client.Get("tasks/"+strconv.Itoa(id), &task, nil); err != nil {
		return task, err
	}
	return task, nil
}

// Create make a new task
func (a *TaskAPI) Create(listID int, title string) (result Task, err error) {
	var task Task

	list, _ := a.client.List.Get(listID)

	createTask := &CreateTask{
		ListID:     listID,
		Title:      title,
		AssigneeID: list.OwnerID,
		Completed:  false,
		Starred:    false,
		DueDate:    fmt.Sprint(time.Now()), // today
	}

	bytes, err := json.Marshal(createTask)
	if err != nil {
		return task, nil
	}

	if err := a.client.Post("tasks", &task, string(bytes)); err != nil {
		return task, err
	}
	return task, nil
}

// Star marks task as starred
func (a *TaskAPI) Star(id int, attach ...bool) (result Task, err error) {
	var task Task

	// check revision number
	existed, err := a.client.Task.Get(id)
	if err != nil {
		return task, err
	}

	starred := true
	if len(attach) > 0 {
		starred = attach[0]
	}

	existed.Starred = starred

	bytes, err := json.Marshal(existed)
	if err != nil {
		return task, err
	}

	if err := a.client.Patch("tasks/"+strconv.Itoa(id), &task, string(bytes)); err != nil {
		return task, err
	}
	return task, nil
}

// UnStar marks new task as unstarred
func (a *TaskAPI) UnStar(id int) (result Task, err error) {
	return a.Star(id, false)
}

// Done marks task as completed
func (a *TaskAPI) Done(id int) (result Task, err error) {
	var task Task

	// check revision number
	existed, err := a.client.Task.Get(id)
	if err != nil {
		return task, err
	}

	existed.Completed = true

	bytes, err := json.Marshal(existed)
	if err != nil {
		return task, err
	}

	if err := a.client.Patch("tasks/"+strconv.Itoa(id), &task, string(bytes)); err != nil {
		return task, err
	}

	return task, nil
}
