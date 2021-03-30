// trying mongoDB crud operations with Go

package main

import (
	"context"
	"employee"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName string
var dbCollection string

func main() {
	client := Connect()
	data := WrapData()
	WriteData(client, data)
}

func init() {
	dbName = "test"
	dbCollection = "blog"
}

func WriteData(client *mongo.Client, data []interface{}) {
	collection := client.Database(dbName).Collection(dbCollection)
	// insertResult, err := collection.InsertMany(context.TODO(), data)
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func WrapData() []interface{} {
	data := []interface{}{ash(), misty(), brock()}
	return data
}

func ash() employee.Personal {
	return employee.Personal{Name: "Bob", Age: 10, Address: "Pallet Town"}
}

func misty() employee.Personal {
	return employee.Personal{Name: "Misty", Age: 10, Address: "Cerulean City"}
}

func brock() employee.Personal {
	return employee.Personal{Name: "Brock", Age: 15, Address: "Pewter City"}
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
