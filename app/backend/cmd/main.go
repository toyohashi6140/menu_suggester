package main

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Text string `json:"greeting"`
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		h := Hello{"Hello, world"}
		res, err := json.Marshal(h)
		if err != nil {
			panic("json error")
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(res)
	})
	http.ListenAndServe(":3030", nil)
}
