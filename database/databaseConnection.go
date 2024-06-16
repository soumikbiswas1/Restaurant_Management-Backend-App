package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBinstance() *mongo.Client {
	 	err := godotenv.Load()
		 if err != nil {
			log.Fatalf("Error loading .env file")
		}
		MongoBb := os.Getenv("MONGODB_URL")
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
