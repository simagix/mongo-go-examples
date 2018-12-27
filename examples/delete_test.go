// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
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
		t.Fatal("delete failed")
	}
}

func TestDeleteMany(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var docs []interface{}
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": bsonx.Int32(1)})
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": bsonx.Int32(2)})
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
		t.Fatal("delete failed, expected", result.DeletedCount)
	}
}
