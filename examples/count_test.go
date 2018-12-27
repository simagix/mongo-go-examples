// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func TestCount(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var count int64
	client = getMongoClient()
	collection = client.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: "color", Value: "Red"}}
	if count, err = collection.Count(ctx, filter); err != nil {
		t.Fatal(err)
	}
	t.Log("count", count)
}
