package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/toyohashi6140/menu_suggester/pkg/handler"
	"github.com/toyohashi6140/menu_suggester/pkg/mongodb"
)

func main() {
	client, err := mongodb.NewClient(os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASSWORD"), "mongo").Connect()
	if err != nil {
		fmt.Println()
		return
	}
	database := mongodb.NewDatabase(client, "menu").SetDB()
	http.HandleFunc("/", handler.Root(database))
	http.HandleFunc("/insert-main-dish", handler.InsertMainDish(database))
	http.ListenAndServe(":3030", nil)
}
