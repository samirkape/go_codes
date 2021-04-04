// Trying JSON encoding with Go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comic struct { // Any field with non-capital first letter will not be marshalled
	Transcript string `json:"transcript"`
}

var URL string = "https://xkcd.com"
var Count int = 1
var dbName string
var dbCollection string

func main() {
	client := initDB()
	fetchWriteJSON(client)
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func fetchWriteJSON(client *mongo.Client) *Comic {
	var result Comic
	for i := 1; i <= Count; i++ {
		_url := URL + "/" + fmt.Sprintf("%d", i) + "/info.0.json"
		resp, err := http.Get(_url)
		if err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil
		}
		writeInDB(client, result)
	}
	return &result
}

func initDB() *mongo.Client {
	dbName = "test"
	dbCollection = "xckdComics"
	return Connect()
}

func writeInDB(client *mongo.Client, data Comic) {
	collection := client.Database(dbName).Collection(dbCollection)
	// insertResult, err := collection.InsertMany(context.TODO(), data)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
}
