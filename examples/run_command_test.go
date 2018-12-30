// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func TestRunCommand(t *testing.T) {
	var err error
	var client *mongo.Client
	var result bson.M
	client = getMongoClient()
	defer client.Disconnect(context.Background())
	command := bson.D{{Key: "isMaster", Value: 1}}
	if err = client.Database("admin").RunCommand(context.Background(), command).Decode(&result); err != nil {
		t.Fatal(err)
	}
	t.Log(stringify(result))
}
