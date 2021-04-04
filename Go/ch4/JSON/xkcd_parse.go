/* Exercise 4.12: The popular web comic xkcd has a JSON interface.
For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites.
Download each URL (once!) and build an offline index.
Write a tool xkcd that, using this index,
    1. prints the URL
    2. transcript of each comic that matches a search term provided on the command line.
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comic struct {
	Transcript string `json:"transcript"`
	ComicNum   int    `json:"ComicNum"`
}

var URL string = "https://xkcd.com"
var URLSuffix = "/info.0.json"
var MongoURI = "mongodb://localhost:27017"
var CollectionName = "xckdComics"
var Count int = 500
var dbName string
var dbCollection string

func main() {
	args := os.Args[1]
	comicNum, _ := strconv.Atoi(args)
	col := InitDB()
	//DeleteCollection(col)
	url, ts := QueryTranscript(col, comicNum)
	fmt.Printf("\tURL\t\t\ttranscript\n")
	fmt.Println(url, ts)
}

func CheckCollection(collection *mongo.Collection) bool {
	names, err := collection.Database().ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return false
	}
	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == CollectionName {
			return true
		}
	}
	return false
}

func DeleteCollection(col *mongo.Collection) {
	_, err := col.DeleteMany(context.TODO(), bson.D{})
	RaiseFatal(err)
}

func QueryTranscript(collection *mongo.Collection, comicNum int) (string, string) {
	input := bson.D{{"comicnum", comicNum}}
	exist := CheckCollection(collection)
	var getDoc Comic
	_url := URL + "/" + fmt.Sprintf("%d", comicNum) + URLSuffix

	if !exist {
		FetchWriteJSON(collection)
	} else {
		log.Printf("The collection exists!.. Fetching from local... ")
		if err := collection.FindOne(context.TODO(), input).Decode(&getDoc); err != nil {
			log.Fatal(err)
		}
	}
	return _url, getDoc.Transcript
}

func Connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	dB := client.Database(dbName).Collection(dbCollection)
	return dB
}

func FetchWriteJSON(collection *mongo.Collection) *Comic {
	var result Comic
	for i := 1; i <= Count; i++ {
		_url := URL + "/" + fmt.Sprintf("%d", i) + URLSuffix
		resp, err := http.Get(_url)
		if err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}

		if resp.StatusCode > 300 {
			continue
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			log.Fatal(err)
			return nil
		}

		result.ComicNum = i
		WriteInDB(collection, result)
	}
	return &result
}

func InitDB() *mongo.Collection {
	dbName = "test"
	dbCollection = "xckdComics"
	return Connect()
}

func WriteInDB(collection *mongo.Collection, insertDoc Comic) {
	// insertResult, err := collection.InsertMany(context.TODO(), data)
	_, err := collection.InsertOne(context.TODO(), insertDoc)
	if err != nil {
		log.Fatal(err)
	}
}

func RaiseFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
