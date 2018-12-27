// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
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
			{$match: { "color": "Red" }},
			{$group: { _id: "$brand", "count": { "$sum": 1 } }},
			{$project: { "brand": "$_id", "_id": 0, "count": 1 }}
		]
	*/

	pipeline := bsonx.Arr{
		bsonx.Document(
			bsonx.Doc{{Key: "$match", Value: bsonx.Document(bsonx.Doc{
				{Key: "color", Value: bsonx.String("Red")},
			})}},
		),
		bsonx.Document(
			bsonx.Doc{{Key: "$group", Value: bsonx.Document(bsonx.Doc{
				{Key: "_id", Value: bsonx.String("$brand")},
				{Key: "count", Value: bsonx.Document(bsonx.Doc{{Key: "$sum", Value: bsonx.Int32(1)}})},
			})}},
		),
		bsonx.Document(
			bsonx.Doc{{
				Key: "$project", Value: bsonx.Document(bsonx.Doc{
					{Key: "brand", Value: bsonx.String("$_id")},
					{Key: "_id", Value: bsonx.Int32(0)},
					{Key: "count", Value: bsonx.Int32(1)},
				}),
			}},
		)}

	t.Log(pipeline)
	collection = client.Database(dbName).Collection(collectionName)
	if cur, err = collection.Aggregate(ctx, pipeline); err != nil {
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
