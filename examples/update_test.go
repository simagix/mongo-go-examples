// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

func TestUpdateOne(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var doc = bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta"}
	var result *mongo.UpdateResult
	client = getMongoClient()
	collection = client.Database(dbName).Collection(collectionName)
	if _, err = collection.InsertOne(ctx, doc); err != nil {
		t.Fatal(err)
	}
	var update bson.M
	json.Unmarshal([]byte(`{ "$set": {"year": 1998}}`), &update)
	if result, err = collection.UpdateOne(ctx, bson.M{"_id": doc["_id"]}, update); err != nil {
		t.Fatal(err)
	}

	if result.ModifiedCount != 1 {
		t.Fatal("update failed, expected 1 but got", result.ModifiedCount)
	}
	collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"})
}

func TestUpdateMany(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var docs []interface{}
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": bsonx.Int32(1)})
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": bsonx.Int32(2)})
	var result *mongo.UpdateResult
	client = getMongoClient()
	collection = client.Database(dbName).Collection(collectionName)
	if _, err = collection.InsertMany(ctx, docs); err != nil {
		t.Fatal(err)
	}
	var update bson.M
	json.Unmarshal([]byte(`{ "$set": {"year": 1998}}`), &update)
	if result, err = collection.UpdateMany(ctx, bson.M{"hometown": "Atlanta"}, update); err != nil {
		t.Fatal(err)
	}
	if result.ModifiedCount != 2 {
		t.Fatal("update failed, expected 2 but got", result.ModifiedCount)
	}
	collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"})
}
