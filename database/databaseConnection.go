package database

import (
	"context"
	"fmt"
	"log"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBinstance() *mongo.Client {
		MongoBb := "mongodb+srv://soumik:0Zlj0locxlzkTkjP@cluster0.d0msb.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	    fmt.Print(MongoBb)

		client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI(MongoBb))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), readpref.Primary())

		if err != nil {
			log.Fatal(err)
		}
	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBinstance()
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
