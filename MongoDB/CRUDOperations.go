package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	collection *mongo.Collection
	client     *mongo.Client
}

//GetConnection - Retunrs the DB client connection and collection
func (db *database) GetConnection(port, dbName, collName string) {

	var err error
	portPath := fmt.Sprintf("mongodb://localhost:%s", port)
	db.client, err = mongo.NewClient(options.Client().ApplyURI(portPath))
	if err != nil {
		log.Fatal(err)
	}

	err = db.client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db.collection = db.client.Database(dbName).Collection(collName)
}

func (db *database) InsertOne(document interface{}) (string, error) {

	res, err := db.collection.InsertOne(context.TODO(), document)
	if err != nil {
		return "", errors.New("Internal error")
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("Internal error while inserting ID")
	}
	return oid.Hex(), nil
}

func (db *database) InsertMany(document []interface{}) ([]interface{}, error) {

	res, err := db.collection.InsertMany(context.TODO(), document)
	if err != nil {
		return nil, errors.New("Internal error")
	}
	return res.InsertedIDs, nil
}

func (db *database) UpdateOne(filter interface{}, update interface{}) error {

	_, err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Internal error")
	}
	return nil
}

func (db *database) UpdateMany(filter interface{}, update interface{}) error {

	_, err := db.collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Internal error")
	}
	return nil
}

func (db *database) DeleteOne(filter interface{}) error {

	_, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Internal error")
	}
	return nil
}

func (db *database) DeleteMany(filter interface{}) error {

	_, err := db.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return errors.New("Internal error")
	}
	return nil
}

func (db *database) Find(filter interface{}) (*mongo.Cursor, error) {

	// Declare Context type object for managing multiple API requests
	cursor, err := db.collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.Background())
	}
	return cursor, err
}

func (db *database) FindOne(filter interface{}) *mongo.SingleResult {

	// Declare Context type object for managing multiple API requests
	return db.collection.FindOne(context.TODO(), filter)
}
