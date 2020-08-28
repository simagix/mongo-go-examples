// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAggregateJSON(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()

	client = getMongoClient()
	defer client.Disconnect(ctx)
	seedCarsData(client, dbName)

	pipeline := `[
		{"$match": { "color": "Red" }},
		{"$group": { "_id": "$brand", "count": { "$sum": 1 } }},
		{"$project": { "brand": "$_id", "_id": 0, "count": 1 }}
	]`
	collection = client.Database(dbName).Collection(collectionName)
	var brands []interface{}
	if brands, err = collection.Distinct(ctx, "brand", bson.D{{Key: "color", Value: "Red"}}); err != nil {
		t.Fatal(err)
	}
	opts := options.Aggregate()
	opts.SetAllowDiskUse(true)
	opts.SetBatchSize(5)
	if cur, err = collection.Aggregate(ctx, MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	if total == 0 || total != len(brands) {
		t.Fatal("expected", len(brands), "but got", total)
	}
}

func TestAggregatePipeline(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()

	client = getMongoClient()
	seedCarsData(client, dbName)

	// this cause warning from go vet
	// pipeline := mongo.Pipeline{
	// 	{{"$match", bson.D{{"color", "Red"}}}},
	// 	{{"$group", bson.D{{"_id", "$brand"}, {"count", bson.D{{"$sum", 1}}}}}},
	// 	{{"$project", bson.D{{"brand", "$_id"}, {"_id", 0}, {"count", 1}}}},
	// }
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "color", Value: "Red"}}}},
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$brand"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}},
		{{Key: "$project", Value: bson.D{{Key: "brand", Value: "$_id"}, {Key: "_id", Value: 0}, {Key: "count", Value: 1}}}},
	}
	collection = client.Database(dbName).Collection(collectionName)
	var brands []interface{}
	if brands, err = collection.Distinct(ctx, "brand", bson.D{{Key: "color", Value: "Red"}}); err != nil {
		t.Fatal(err)
	}
	opts := options.Aggregate()
	opts.SetAllowDiskUse(true)
	opts.SetBatchSize(5)
	if cur, err = collection.Aggregate(ctx, pipeline, opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	if total == 0 || total != len(brands) {
		t.Fatal("expected", len(brands), "but got", total)
	}
}

func TestAggregateBSOND(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()
	size := 10
	client = getMongoClient()
	seedCarsData(client, dbName)
	pipeline := []bson.D{bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: size}}}}}
	collection = client.Database(dbName).Collection(collectionName)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, pipeline, opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	if total == 0 {
		t.Fatal("expected ", size, "but got", total)
	}
}
