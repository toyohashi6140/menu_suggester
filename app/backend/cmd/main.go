package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/toyohashi6140/menu_suggester/pkg/mongodb"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		mongodb := mongodb.New(os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASSWORD"), "mongo").SetDB("menu").SetCollection("maindish")
		collection, err := mongodb.Connect()
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
			rw.Write(res)
			return
		}
		res, err := json.Marshal(map[string]interface{}{"mongo": collection})
		if err != nil {
			panic("json error")
		}

		rw.Write(res)
	})
	http.ListenAndServe(":3030", nil)
}
