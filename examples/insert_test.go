// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func TestInsertOne(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var doc = bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta"}
	var result *mongo.InsertOneResult
	client = getMongoClient()
	defer client.Disconnect(ctx)
	collection = client.Database(dbName).Collection(collectionName)
	if result, err = collection.InsertOne(ctx, doc); err != nil {
		t.Fatal(err)
	}
	if result.InsertedID != doc["_id"] {
		t.Fatal(result.InsertedID, doc["_id"])
	}
	collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"})
}

// TestInsertOneWithOptions insert with w: "majority"
func TestInsertOneWithOptions(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var doc = bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta"}
	var result *mongo.InsertOneResult
	client = getMongoClient()
	defer client.Disconnect(ctx)
	opts := options.Collection()
	opts.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	collection = client.Database(dbName).Collection(collectionName, opts)
	if result, err = collection.InsertOne(ctx, doc); err != nil {
		t.Fatal(err)
	}
	if result.InsertedID != doc["_id"] {
		t.Fatal(result.InsertedID, doc["_id"])
	}
	collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"})
}

func TestInsertMany(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var docs []interface{}
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": 1})
	docs = append(docs, bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta", "counter": 2})
	var result *mongo.InsertManyResult
	client = getMongoClient()
	defer client.Disconnect(ctx)
	collection = client.Database(dbName).Collection(collectionName)
	if result, err = collection.InsertMany(ctx, docs); err != nil {
		t.Fatal(err)
	}
	for _, doc := range docs {
		isFound := false
		for _, id := range result.InsertedIDs {
			if id == doc.(bson.M)["_id"] {
				isFound = true
				continue
			}
		}
		if !isFound {
			t.Fatal(doc.(bson.M)["_id"], "not inserted")
		}
	}
	collection.DeleteMany(ctx, bson.M{"hometown": "Atlanta"})
}
