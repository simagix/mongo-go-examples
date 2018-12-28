// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/simagix/keyhole/mdb"
)

func TestAggregate(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M

	client = getMongoClient()
	seedCarsData(client, dbName)

	pipeline := `[
		{"$match": { "color": "Red" }},
		{"$group": { "_id": "$brand", "count": { "$sum": 1 } }},
		{"$project": { "brand": "$_id", "_id": 0, "count": 1 }}
	]`
	collection = client.Database(dbName).Collection(collectionName)
	opts := options.Aggregate()
	opts.SetAllowDiskUse(true)
	opts.SetBatchSize(5)
	if cur, err = collection.Aggregate(ctx, mdb.GetAggregatePipeline(pipeline), opts); err != nil {
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
