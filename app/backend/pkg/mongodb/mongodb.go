package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	user, password, host, db, collection string
}

func New(user, password, host string) *mongodb {
	return &mongodb{
		user:     user,
		password: password,
		host:     host,
	}
}

// SetDB set database name
func (m *mongodb) SetDB(db string) *mongodb {
	m.db = db
	return m
}

// SetCollection set collection name
func (m *mongodb) SetCollection(collection string) *mongodb {
	m.collection = collection
	return m
}

// Connect connect mongoDB client and get collection, then return *mongo.collection
func (m *mongodb) Connect() (*mongo.Collection, error) {
	if m.db == "" {
		return nil, errors.New("no database selected")
	}
	if m.collection == "" {
		return nil, errors.New("no collection selected")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017", m.user, m.password, m.host)),
	)
	if err != nil {
		return nil, err
	}
	collection := client.Database(m.db).Collection(m.collection)
	return collection, nil
}
