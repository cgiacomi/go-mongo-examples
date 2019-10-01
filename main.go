package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findOneAndUpdate(url string, db string, coll string) (bson.M, error) {

	// 1) Create the context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2) Create the connection
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err.Error())
	}

	// 2) Choose the database
	database := conn.Database(db)

	// 4) Set the collection
	collection := database.Collection(coll)

	// 5) Create the search filter
	filter := bson.M{"name": "luke"}

	// 6) Create the update
	update := bson.M{
		"$set": bson.M{"lastname": "skywalker"},
	}

	// 7) Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// 8) Find one result and update it
	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}

	// 9) Decode the result
	doc := bson.M{}
	decodeErr := result.Decode(&doc)

	return doc, decodeErr
}

func main() {
	url := "mongodburl"
	db := "jedi"
	coll := "characters"

	doc, err := findOneAndUpdate(url, db, coll)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Printf("%+v", doc)
}
