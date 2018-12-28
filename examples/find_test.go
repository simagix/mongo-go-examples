// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func TestFindOne(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var doc bson.M
	client = getMongoClient()
	seedCarsData(client, dbName)
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: "color", Value: "Red"}}
	if err = collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		t.Fatal(err)
	}
	t.Log(doc["brand"])
}

func TestFindMany(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M
	client = getMongoClient()
	seedCarsData(client, dbName)
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: "color", Value: "Red"}}
	if cur, err = collection.Find(ctx, filter); err != nil {
		t.Fatal(err)
	}
	total := 0
	for cur.Next(ctx) {
		cur.Decode(&doc)
		total++
	}
	t.Log("total", total)
}

func TestFindManyWithOptions(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M
	client = getMongoClient()
	seedCarsData(client, dbName)
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{}
	opts := options.Find()
	opts.SetBatchSize(int32(10))
	opts.SetLimit(int64(2))
	opts.SetSkip(int64(10))
	opts.SetProjection(bson.M{"_id": 0, "filters": 0})
	if cur, err = collection.Find(ctx, filter, opts); err != nil {
		t.Fatal(err)
	}
	total := 0
	for cur.Next(ctx) {
		cur.Decode(&doc)
		t.Log(doc)
		total++
	}
	if total != 2 {
		t.Fatal("find failed, expected 2 but got", total)
	}
}
