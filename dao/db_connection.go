package dao

import (
	"context"
	"log"
	"time"

	// "github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}
	clientdb := client.Database("Wall-E")
	return clientdb

}
