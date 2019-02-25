// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCountDocuments(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var count int64
	client = getMongoClient()
	seedCarsData(client, dbName)
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: "color", Value: "Red"}}
	if count, err = collection.CountDocuments(ctx, filter); err != nil {
		t.Fatal(err)
	}
	t.Log("count", count)
}
