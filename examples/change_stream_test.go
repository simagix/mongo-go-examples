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

var collection = "oplogs"

// example: argos "mongodb://localhost:27017/argos?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'

func TestChangeStreamClient(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
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
	c := client.Database(cs.Database).Collection(collection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
		client.Disconnect(context.Background())
	}(c)

	stream := NewChangeStream()
	stream.SetPipeline(pipeline)
	// stream.Watch(client)
}

func TestChangeStreamDatabase(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
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
	c := client.Database(cs.Database).Collection(collection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
		client.Database(cs.Database).Drop(context.Background())
	}(c)

	stream := NewChangeStream()
	stream.SetDatabase(cs.Database)
	stream.SetPipeline(pipeline)
	stream.Watch(client)
}

func TestChangeStreamCollection(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
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
	c := client.Database(cs.Database).Collection(collection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
	}(c)

	stream := NewChangeStream()
	stream.SetCollection(collection)
	stream.SetDatabase(cs.Database)
	stream.SetPipeline(pipeline)
	stream.Watch(client)
}

func TestChangeStreamCollectionWithPipeline(t *testing.T) {
	var err error
	var client *mongo.Client
	var cs connstring.ConnString
	var ctx = context.Background()
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
	var pipeline = mdb.GetAggregatePipeline(`[{"$match": {"operationType": {"$in": ["update", "delete"] } }}]`)
	c := client.Database(cs.Database).Collection(collection)
	c.InsertOne(ctx, bson.M{"city": "Atlanta"})

	go func(c *mongo.Collection) {
		execute(c)
	}(c)

	stream := NewChangeStream()
	stream.SetCollection(collection)
	stream.SetDatabase(cs.Database)
	stream.SetPipeline(pipeline)
	stream.Watch(client)
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
