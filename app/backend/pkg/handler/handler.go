package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/toyohashi6140/menu_suggester/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler func(rw http.ResponseWriter, r *http.Request)

func Root(d *mongo.Database) Handler {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		collection, err := mongodb.NewCollection(d, "mainDish").SetCollection()
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
	}
}

func InsertMainDish(d *mongo.Database) Handler {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		collection, err := mongodb.NewCollection(d, "mainDish").SetCollection()
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
			rw.Write(res)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		j, err := ioutil.ReadFile(filepath.Join("../", "mainDish.js"))
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
			rw.Write(res)
			return
		}
		result, err := collection.InsertOne(ctx, j)
		id := result.InsertedID.(string)
		rw.Write([]byte(id))
	}
}
