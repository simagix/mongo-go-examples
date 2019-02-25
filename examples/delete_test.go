// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteOne(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var doc = bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta"}
	var result *mongo.DeleteResult
	client = getMongoClient()
	collection = client.Database(dbName).Collection(collectionName)
	if _, err = collection.InsertOne(ctx, doc); err != nil {
		t.Fatal(err)
	}
	if result, err = collection.DeleteOne(ctx, bson.M{"_id": doc["_id"]}); err != nil {
		t.Fatal(err)
	}

	if result.DeletedCount != 1 {
		t.Fatal("delete failed, expected 1 but got", result.DeletedCount)
	}
}

func TestDeleteMany(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var docs []interface{}
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": 1})
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": 2})
	var result *mongo.DeleteResult
	client = getMongoClient()
	collection = client.Database(dbName).Collection(collectionName)
	if _, err = collection.InsertMany(ctx, docs); err != nil {
		t.Fatal(err)
	}
	if result, err = collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"}); err != nil {
		t.Fatal(err)
	}
	if result.DeletedCount != 2 {
		t.Error("delete failed, expected 2 but got", result.DeletedCount)
	}
}
