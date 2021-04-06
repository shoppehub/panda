package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Resource struct {
	DB     *mongo.Database
	Client *mongo.Client
}

var Instance Resource

// 当前连接的数据库实例
var DB *mongo.Database

func New() (*Resource, error) {

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		return nil, err
	}
	// fmt.Println(client)
	db := client.Database(os.Getenv("MONGO_DBNAME"))
	DB = db
	Instance = Resource{DB: db, Client: client}

	return &Instance, nil

}

// Close closes the mongo-go-driver connection.
func (d *Resource) Close() {
	d.Client.Disconnect(context.Background())
}
