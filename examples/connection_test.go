// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"github.com/simagix/keyhole/mdb"
	"github.com/simagix/keyhole/sim"
)

const dbName = "keyhole"
const collectionName = "cars"

func getMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client

	uri := "mongodb://localhost/keyhole?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}

	if client, err = mdb.NewMongoClient(uri); err != nil {
		panic(err)
	}
	cs, _ := connstring.Parse(uri)
	f := sim.NewFeeder()
	f.SetTotal(100)
	f.SetIsDrop(true)
	f.SetDatabase(cs.Database)
	f.SetShowProgress(false)
	f.SeedCars(client)
	return client
}
