// Copyright 2019 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDistinct(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()
	client = getMongoClient()
	defer client.Disconnect(ctx)
	seedCarsData(client, dbName)
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{{}}

	if cur, err = collection.Find(ctx, filter); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	amap := map[string]string{}
	var doc bson.M

	for cur.Next(ctx) {
		cur.Decode(&doc)
		brand := doc["brand"].(string)
		amap[brand] = brand
	}

	var brands []interface{}
	if brands, err = collection.Distinct(ctx, "brand", filter); err != nil {
		t.Fatal(err)
	}
	if len(brands) != len(amap) {
		t.Fatal("find failed, expected", len(amap), "but got", len(brands))
	}
}
