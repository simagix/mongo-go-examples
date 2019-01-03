// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/simagix/keyhole/mdb"
)

/*
 * count vehicles by style and display all brands and a total count of each style
 */
func TestAggregateGroup(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M

	client = getMongoClient()
	total := seedCarsData(client, dbName)

	pipeline := `
	[{
		"$group": {
			"_id": "$style",
			"brand": {
				"$addToSet": "$brand"
			},
			"count": {
				"$sum": 1
			}
		}
	}]`
	collection = client.Database(dbName).Collection(collectionName)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, mdb.MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	count := int64(0)
	for cur.Next(ctx) {
		cur.Decode(&doc)
		t.Log(doc["_id"], doc["count"])
		count += int64(doc["count"].(float64))
	}

	if total != count {
		t.Fatal("expected", total, "but got", count)
	}
}
