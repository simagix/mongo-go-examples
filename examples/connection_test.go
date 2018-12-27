// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/keyhole/mdb"
)

const dbName = "keyhole"
const collectionName = "cars"

func getMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client

	if client, err = mdb.NewMongoClient("mongodb://localhost/keyhole?replicaSet=replset"); err != nil {
		panic(err)
	}

	return client
}
