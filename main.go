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

	pp.Print(client.Get("lists"))
}
