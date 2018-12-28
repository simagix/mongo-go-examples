// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"github.com/simagix/keyhole/mdb"
)

// example: argos "mongodb://localhost:27017/argos?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'
func TestChangeStream(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
	var colection = "oplogs"
	var uri = "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	if cs, err = connstring.Parse(uri); err != nil {
		t.Fatal(err)
	}
	if client, err = mdb.NewMongoClient(uri); err != nil {
		t.Fatal(err)
	}
	var pipeline []bson.D
	pipeline = mongo.Pipeline{}
	c := client.Database(cs.Database).Collection(colection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
	}(c)

	ChangeStream(client, cs.Database, colection, pipeline)
}

func TestChangeStreamWithPipeline(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
	var colection = "oplogs"
	var uri = "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	if cs, err = connstring.Parse(uri); err != nil {
		t.Fatal(err)
	}
	if client, err = mdb.NewMongoClient(uri); err != nil {
		t.Fatal(err)
	}
	var pipeline = mdb.GetAggregatePipeline(`[{"$match": {"operationType": "update"}}]`)
	c := client.Database(cs.Database).Collection(colection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
	}(c)

	ChangeStream(client, cs.Database, colection, pipeline)
}

func execute(c *mongo.Collection) {
	time.Sleep(3 * time.Second) // wait for change stream to init
	var doc = bson.M{"_id": primitive.NewObjectID(), "hometown": "Atlanta"}
	c.InsertOne(context.Background(), doc)
	var update bson.M
	json.Unmarshal([]byte(`{ "$set": {"year": 1998}}`), &update)
	c.UpdateOne(context.Background(), bson.M{"_id": doc["_id"]}, update)
	c.DeleteMany(context.Background(), bson.M{"hometown": "Atlanta"})
	time.Sleep(3 * time.Second) // wait for CS to print messages
	c.Drop(context.Background())
}
