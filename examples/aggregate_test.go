// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func TestAggregate(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M

	client = getMongoClient()

	/*
		[
			{"$match": { "color": "Red" }},
			{"$group": { _id: "$brand", "count": { "$sum": 1 } }},
			{"$project": { "brand": "$_id", "_id": 0, "count": 1 }}
		]
	*/

	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "color", Value: "Red"}}}},
		bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$brand"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}},
		bson.D{{Key: "$project", Value: bson.D{{Key: "brand", Value: "$_id"},
			{Key: "_id", Value: 0},
			{Key: "count", Value: 1}}}},
	}

	collection = client.Database(dbName).Collection(collectionName)
	opts := options.Aggregate()
	opts.SetAllowDiskUse(true)
	opts.SetBatchSize(5)
	if cur, err = collection.Aggregate(ctx, pipeline, opts); err != nil {
		t.Fatal(err)
	}
	total := 0
	for cur.Next(ctx) {
		cur.Decode(&doc)
		t.Log(doc)
		total++
	}
	t.Log("total", total)
}
