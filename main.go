package main

import (
	"github.com/k0kubun/pp"
	"github.com/qube81/wunderlist-api-go/wunderlist"
	"os"
)

func main() {

	clientID := os.Getenv("WL_CLIENT_ID")
	clientSecret := os.Getenv("WL_ACCESS_TOKEN")

	client := wunderlist.NewClient(clientID, clientSecret)

	lists, _ := client.List.GetAll()
	list, _ := client.List.Get(lists[1].ID)
	user, _ := client.User.Get()
	tasks, _ := client.Task.GetByListID(lists[1].ID, false)
	task, _ := client.Task.Get(tasks[0].ID)

	pp.Print(lists)
	pp.Print(list)
	pp.Print(user)
	pp.Print(tasks)
	pp.Print(task)

}
