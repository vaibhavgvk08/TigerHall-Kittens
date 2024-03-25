package database

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoDB implements iDataBase
type MongoDB struct {
	dbclient  *mongo.Client
	dbcontext context.Context
	cancel    context.CancelFunc
}

func Connect(uri string) *MongoDB {
	mongodb := &MongoDB{}
	ctx, cancelFunc := context.WithTimeout(context.Background(),
		1000*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	mongodb.dbclient = client
	mongodb.dbcontext = ctx
	mongodb.cancel = cancelFunc
	return mongodb
}

func (obj *MongoDB) Find(dataBase, col string, query interface{}) (result []byte, err error) { //*mongo.Cursor
	collection := obj.dbclient.Database(dataBase).Collection(col)
	cursor, err := collection.Find(obj.dbcontext, query)
	var jsonBytes []byte

	for cursor.Next(context.Background()) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Fatal(err)
		}

		// Marshal document to JSON bytes
		docBytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal(err)
		}
		if len(jsonBytes) > 0 {
			jsonBytes = append(jsonBytes, ',')
		}
		// Append JSON bytes to the slice
		jsonBytes = append(jsonBytes, docBytes...)
	}
	jsonBytes = append([]byte{'['}, jsonBytes...)
	jsonBytes = append(jsonBytes, ']')

	return jsonBytes, err
}

func (obj *MongoDB) Insert(dataBase, col string, doc interface{}) (interface{}, error) { //*mongo.InsertOneResult
	collection := obj.dbclient.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(obj.dbcontext, doc)
	fmt.Println(doc, result, err)
	return result.InsertedID, err
}

func (obj *MongoDB) Update(dataBase, col string, filter, update interface{}) (result interface{}, err error) { //*mongo.UpdateResult

	collection := obj.dbclient.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(obj.dbcontext, filter, update)
	return
}

func (obj *MongoDB) Delete(dataBase, col string, query interface{}) (result interface{}, err error) { //*mongo.DeleteResult

	collection := obj.dbclient.Database(dataBase).Collection(col)
	result, err = collection.DeleteOne(obj.dbcontext, query)
	return

}

func (obj *MongoDB) Close() {
	defer obj.cancel()
	defer func() {
		if err := obj.dbclient.Disconnect(obj.dbcontext); err != nil {
			panic(err)
		}
	}()
}
