package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	User    *mongo.Collection
	Url     *mongo.Collection
	Tracker *mongo.Collection
}

var CurrentDb *DB

func CreateDBInstance() {
	connectionString := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	userCollName := os.Getenv("DB_USER_COLLECTION_NAME")
	urlCollName := os.Getenv("DB_URL_COLLECTION_NAME")
	trackerCollName := os.Getenv("DB_TRACKER_COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
		return
	}

	fmt.Println("Connected to MongoDB")

	userCollection := client.Database(dbName).Collection(userCollName)
	urlCollection := client.Database(dbName).Collection(urlCollName)
	trackerCollection := client.Database(dbName).Collection(trackerCollName)

	CurrentDb = &DB{User: userCollection, Url: urlCollection, Tracker: trackerCollection}
}
