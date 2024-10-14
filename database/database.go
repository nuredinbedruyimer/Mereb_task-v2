package database

import (
	"Mereb-V2/constants"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancell := context.WithTimeout(context.Background(), constants.TIME_OUT)

	defer cancell()

	client, connectingError := mongo.Connect(ctx, clientOptions)

	if connectingError != nil {
		fmt.Println("Error When Connectiong With MongoDB : ", connectingError)
		return nil
	}

	err := client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Error When Ping Our Client : ", err)
		return nil
	}
	return client

}

var Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection = client.Database("mereb-task").Collection(collectionName)

	return collection
}
