package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

func NewMongoConnect() (*mongo.Database, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("MONGODB_DATABASE"))

	return db, err
}
