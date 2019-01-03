// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/keyhole/sim"
)

const dbName = "argos"
const collectionName = "cars"
const collectionFavorites = "favorites"
const collectionExamples = "examples"

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

// SeedCarsData wraps seedCarsData
func SeedCarsData(client *mongo.Client, database string) int64 {
	return seedCarsData(client, database)
}

func seedCarsData(client *mongo.Client, database string) int64 {
	var err error
	var count int64
	collection := client.Database(dbName).Collection(collectionName)
	filter := bson.D{{}}
	if count, err = collection.Count(context.Background(), filter); err != nil {
		fmt.Println("===>", err)
		return 0
	}
	if count == 0 {
		f := sim.NewFeeder()
		f.SetTotal(100)
		f.SetIsDrop(true)
		f.SetDatabase(database)
		f.SetShowProgress(false)
		f.SeedCars(client)
		return int64(100)
	}
	return count
}

func seedFavoritesData(client *mongo.Client, database string) int64 {
	var err error
	var count int64
	collection := client.Database(dbName).Collection(collectionFavorites)
	filter := bson.D{{}}
	if count, err = collection.Count(context.Background(), filter); err != nil {
		fmt.Println(err)
		return 0
	}
	if count == 0 {
		f := sim.NewFeeder()
		f.SetTotal(100)
		f.SetIsDrop(true)
		f.SetDatabase(database)
		f.SetShowProgress(false)
		f.SeedFavorites(client)
		return int64(100)
	}
	return count
}

func stringify(doc interface{}, opts ...string) string {
	if len(opts) == 2 {
		b, _ := json.MarshalIndent(doc, opts[0], opts[1])
		return string(b)
	}
	b, _ := json.Marshal(doc)
	return string(b)
}
