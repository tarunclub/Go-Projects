package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://tarun:J@v@scr!ptes6@cluster0.s2vx0mc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect to mongo

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reference is ready")
}
