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

	allLists, _ := client.List.GetAll()
	/*
		lists, _ := client.List.Get(allLists[1].ID)
		user, _ := client.User.Get()
	*/

	task, _ := client.Task.GetByListID(allLists[0].ID, false)

	pp.Print(task)

	/*
		pp.Print(lists)
		pp.Print(user)
	*/
}
