// Copyright 2019 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnectionPool(t *testing.T) {
	t.Log("TestConnectionPool")
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	if client, err = getMongoClient(); err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)
	seedCarsData(client, dbName)
	channel := make(chan string)
	nThreads := 3
	for i := 0; i < nThreads; i++ {
		go func(client *mongo.Client, channel chan string, i int) {
			var doc bson.M
			collection = client.Database(dbName).Collection(collectionName)
			filter := bson.D{{Key: "color", Value: "Red"}}
			if err = collection.FindOne(ctx, filter).Decode(&doc); err != nil {
				t.Fatal(err)
			}
			channel <- fmt.Sprintf("child %d completed", i)
		}(client, channel, i)
	}
	cnt := 0
	for {
		msg := <-channel
		t.Log(msg)
		cnt++
		if cnt == nThreads {
			break
		}
		time.Sleep(time.Second * 1)
	}
}
