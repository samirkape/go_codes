//  Trying mongo Go driver

package main

import (
	"context"
	"fmt"
	"log"

	"employee"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	exp_push(client)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func exp_push(client *mongo.Client) {
	collection := client.Database("test").Collection("blog")
	fmt.Println(collection)
	ash := employee.Personal{Name: "Bob", Age: 10, Address: "Pallet Town"}
	misty := employee.Personal{Name: "Misty", Age: 10, Address: "Cerulean City"}
	brock := employee.Personal{Name: "Brock", Age: 15, Address: "Pewter City"}
	trainers := []interface{}{ash, misty, brock}
	insertResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
	// InsertResult, err := collection.InsertOne()

}
