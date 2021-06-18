// mparser-db is responsible for handling database related operation
// which may include connect, write, query
package mybot

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WriteData uses mongodb's  InsertMany()  function to insert documents to a
// dbName database and CollectionName collection
func WriteData(client *mongo.Client, DbName string, CollectionName string, data []interface{}) *mongo.Collection {
	//Create a handle to the respective collection in the database.
	collection := client.Database(DbName).Collection(CollectionName)
	//Perform InsertMany operation & validate against the error.
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	return collection
}

func RemoveDuplicates(client *mongo.Client, DB string) {
	collections := ListCollections(client, DB)
	for _, coll := range collections {
		FindDoc(client, DB, coll, "")
	}
}

func ListCollections(client *mongo.Client, DB string) []string {
	collections, err := client.Database(DB).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return collections
}

func FindDoc(client *mongo.Client, DB string, Collection string, Search string) ([]SplitLink, error) {
	var LinkList []SplitLink
	var cur *mongo.Cursor
	var findError error

	collection := client.Database(DB).Collection(Collection)

	//Define filter query for fetching specific document from collection
	if Search != "" {
		filter := bson.M{"info": "/" + Search + "/"}
		cur, findError = collection.Find(context.TODO(), filter)
	} else {
		filter := bson.D{} //bson.D{{}} specifies 'all documents'
		cur, findError = collection.Find(context.TODO(), filter)
	}
	//Create a handle to the respective collection in the database.
	//Perform Find operation & validate against the error.
	if findError != nil {
		return nil, findError
	}
	defer cur.Close(context.TODO())
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := SplitLink{}
		err := cur.Decode(&t)
		if err != nil {
			return nil, err
		} else {
			LinkList = append(LinkList, t)
		}
	}
	return LinkList, nil
}

func FindDeleteDoc(client *mongo.Client, DB string, Collection string) error {
	//Define filter query for fetching specific document from collection
	filter := bson.D{} //bson.D{{}} specifies 'all documents'
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Collection)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return findError
	}
	defer cur.Close(context.TODO())
	namemap := make(map[string]struct{})
	var estruct struct{}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := SplitLink{}
		err := cur.Decode(&t)
		if err != nil {
			return err
		}
		// if key is already present, then its a duplicate, remove it
		if _, ok := namemap[t.URL]; ok {
			DeleteOne(client, DB, Collection, t.ID)
		} else {
			namemap[t.URL] = estruct
		}
	}
	// once exhausted, close the cursor
	return nil
}

func DeleteOne(client *mongo.Client, DB string, Collection string, ID primitive.ObjectID) error {
	filter := bson.M{"_id": ID}

	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Collection)
	//Perform DeleteOne operation & validate against the error.
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

// DbConnect establish connection to mongodb cloud database for a given URI and
// returns *mongo.Client  which needs to be used for further operations on database.
func GetDbClient() *mongo.Client {
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

// DbWrite takes final meta as an input which is
// name, url, short info passed at once for a particular collection.
// it does some pre-processing and calls WriteData for actual insertion.
func DbWritePkgs(final []Package, client *mongo.Client, DbName string) {
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

func WriteUser(client *mongo.Client, DbName string, CollectionName string, data interface{}) *mongo.Collection {
	//Create a handle to the respective collection in the database.
	collection := client.Database(DbName).Collection(CollectionName)
	//Perform InsertMany operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	return collection
}
