// mparser-db is responsible for handling database related operation
// which may include connect, write, query
package mparser

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DbName = "mparserdb"

// DbWrite takes final meta as an input which is
// name, url, short info passed at once for a particular collection.
// it does some pre-processing and calls WriteData for actual insertion.
func DbWritePkgList(final []Package, client *mongo.Client, DbName string) {
	for i := 0; i < len(final); i++ {
		e := final[i]
		var data []interface{}
		title := e.Details.Title
		for j := 0; j < len(e.Details.Line.LinkDetails); j++ {
			data = append(data, e.Details.Line.LinkDetails[j])
		}
		WriteData(client, DbName, title, data)
	}
}

// WriteData uses mongodb's  InsertMany()  function to insert documents to a
// dbName database and CollectionName collection
func WriteData(client *mongo.Client, DbName string, CollectionName string, data []interface{}) *mongo.Collection {
	collection := client.Database(DbName).Collection(CollectionName)
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	return collection
}

// func Query(client *mongo.Client, CollectionName string, key string) {
// 	collection := client.Database(dbName).Collection(CollectionName)
// 	Info := SplitLink{}
// 	if err := collection.Find(context.Background(), bson.D{key}).Decode(&Info); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(Info)
// }

// DbConnect establish connection to mongodb cloud database for a given URI and
// returns *mongo.Client  which needs to be used for further operations on database.
func DbConnect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
