// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
 * https://docs.mongodb.com/manual/reference/operator/aggregation/graphLookup/
 */
func TestAggregateGraphLookup(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()
	var doc bson.M

	if client, err = getMongoClient(); err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)
	seedCarsData(client, dbName)

	pipeline := `
	[{
		"$graphLookup": {
			"from": "employees",
			"startWith": "$manager",
			"connectFromField": "manager",
			"connectToField": "_id",
			"as": "employeeHierarchy"
		}
	}]
	`

	collection = client.Database(dbName).Collection("employees")
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	count := int64(0)
	for cur.Next(ctx) {
		cur.Decode(&doc)
		count++
	}

	if 0 == count {
		t.Fatal("no doc found")
	}
}
