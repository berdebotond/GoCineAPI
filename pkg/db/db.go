package database

import (
	"context"
	config "interview_test_KenTech/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	log.Println("Connecting to the Mongodb")
	mongodb_options := options.Client().ApplyURI(config.CFG.MongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, mongodb_options)
	if err != nil {
		log.Fatal("Failed to Connect to the Mongodb" + err.Error())
	}
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil
	}
	// list collection names
	collections, err := client.Database("Cluster0").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println(collections)
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(config.CFG.DbName).Collection(CollectionName)
	return collection
}

// Add this function in db.go
func FilmData(client *mongo.Client, CollectionName string) *mongo.Collection {
	return client.Database(config.CFG.DbName).Collection(CollectionName)
}
