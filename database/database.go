package database

import (
	"fmt"
	"log"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

type Todo struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Todo string  `json:"todo,omitempty" bson:"todo,omitempty"`
}

var Client *mongo.Client

func Db() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MongoUrl")))

	err := Client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	databases, err := Client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
}