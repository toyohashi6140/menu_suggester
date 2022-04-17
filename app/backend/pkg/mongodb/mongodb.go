package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
)

type mongodb struct {
	credential           options.Credential
	host, db, collection string
}

func New(user, password, host string) *mongodb {
	return &mongodb{
		credential: options.Credential{
			AuthMechanism: auth.SCRAMSHA256,
			Username:      user,
			Password:      password,
		},
		host: host,
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", m.host)).SetAuth(m.credential),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	collection := client.Database(m.db).Collection(m.collection)
	return collection, nil
}
