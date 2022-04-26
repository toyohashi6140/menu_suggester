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

type (
	client struct {
		credential options.Credential
		host       string
	}

	database struct {
		client *mongo.Client
		db     string
	}

	collection struct {
		database   *mongo.Database
		collection string
	}
)

func NewClient(user, password, host string) *client {
	return &client{
		credential: options.Credential{
			AuthMechanism: auth.SCRAMSHA256,
			Username:      user,
			Password:      password,
		},
		host: host,
	}
}

// Connect connect mongoDB client and get collection, then return *mongo.collection
func (c *client) Connect() (*mongo.Client, error) {
	if c.host == "" {
		return nil, errors.New("no host setting")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", c.host)).SetAuth(c.credential),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

func NewDatabase(c *mongo.Client, d string) *database {
	return &database{client: c, db: d}
}

func (d *database) SetDB() *mongo.Database {
	if d.db == "" {
		d.db = "menu"
	}
	return d.client.Database(d.db)
}

func NewCollection(d *mongo.Database, c string) *collection {
	return &collection{database: d, collection: c}
}

func (c *collection) SetCollection() (*mongo.Collection, error) {
	if c.collection == "" {
		return nil, errors.New("no collection selected")
	}
	return c.database.Collection(c.collection), nil
}
