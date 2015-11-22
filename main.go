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
	list, _ := client.List.Get(allLists[1].ID)

	pp.Print(allLists)
	pp.Print(list)
}
