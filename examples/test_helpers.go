// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/keyhole/sim"
)

const dbName = "argos"
const collectionName = "cars"
const collectionFavorites = "favorites"

func getMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client

	uri := "mongodb://localhost/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}

	if client, err = mongo.Connect(context.Background(), uri); err != nil {
		panic(err)
	}
	return client
}

func seedCarsData(client *mongo.Client, database string) {
	var err error
	var count int64
	collection := client.Database(dbName).Collection(collectionName)
	if count, err = collection.Count(context.Background(), nil); err != nil {
		fmt.Println(err)
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

func seedFavoritesData(client *mongo.Client, database string) {
	var err error
	var count int64
	collection := client.Database(dbName).Collection(collectionFavorites)
	if count, err = collection.Count(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}
	if count == 0 {
		f := sim.NewFeeder()
		f.SetTotal(100)
		f.SetIsDrop(true)
		f.SetDatabase(database)
		f.SetShowProgress(false)
		f.SeedFavorites(client)
	}
}
