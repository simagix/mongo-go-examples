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
	if doc["color"] != "Red" {
		t.Fatal("not matched", doc["color"])
	}
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
	count, _ := collection.Count(ctx, filter)
	if cur, err = collection.Find(ctx, filter); err != nil {
		t.Fatal(err)
	}
	total := int64(0)
	for cur.Next(ctx) {
		cur.Decode(&doc)
		total++
	}
	if total != count {
		t.Fatal("find failed, expected", count, "but got", total)
	}
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
	limit := 3
	filter := bson.D{}

	// set options
	opts := options.Find()
	opts.SetBatchSize(int32(10))
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(20))
	opts.SetProjection(bson.M{"_id": 0, "filters": 0})
	opts.SetSort(bson.D{{Key: "brand", Value: 1}, {Key: "style", Value: -1}})
	if cur, err = collection.Find(ctx, filter, opts); err != nil {
		t.Fatal(err)
	}
	total := 0
	for cur.Next(ctx) {
		cur.Decode(&doc)
		t.Log(doc["brand"], doc["style"], doc["year"])
		total++
	}
	if total != limit {
		t.Fatal("find failed, expected", limit, "but got", total)
	}
}
