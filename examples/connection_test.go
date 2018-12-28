// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"os"

	"github.com/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/keyhole/mdb"
	"github.com/simagix/keyhole/sim"
)

const dbName = "argos"
const collectionName = "cars"

func getMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client

	uri := "mongodb://localhost/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}

	if client, err = mdb.NewMongoClient(uri); err != nil {
		panic(err)
	}
	return client
}

func seedCarsData(client *mongo.Client, database string) {
	var err error
	var count int64
	collection := client.Database(dbName).Collection(collectionName)
	filter := bson.D{}
	if count, err = collection.Count(context.Background(), filter); err != nil {
		return
	}
	if count == 0 {
		f := sim.NewFeeder()
		f.SetTotal(100)
		f.SetIsDrop(true)
		f.SetDatabase(database)
		f.SetShowProgress(false)
		f.SeedCars(client)
	}
}
