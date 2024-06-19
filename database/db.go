package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB             *mongo.Database
	UserCollection *mongo.Collection
)

func InitDatabase() {
	fmt.Println("Connecting to mongo cluster")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}

	//Ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to mongo cluster")
	DATABASE = client.Database(DATABASE_NAME)
}

var DATABASE *mongo.Database

const DATABASE_URL = "mongodb+srv://cgaas:rvyuMzkZXfLp52m7@cgaas.bbsin5h.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME = "Deal-APP234"
