// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/simagix/keyhole/sim"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "argos"
const collectionName = "cars"
const collectionFavorites = "favorites"
const collectionExamples = "examples"

var once sync.Once
var (
	instance *mongo.Client
)

func connectSingleton() *mongo.Client {
	if instance == nil {
		instance = getMongoClient()
	}
	return instance
}

func getMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client
	uri := "mongodb://localhost/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	client.Ping(context.Background(), nil)
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
	if count, err = collection.CountDocuments(context.Background(), filter); err != nil {
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
	if count, err = collection.CountDocuments(context.Background(), filter); err != nil {
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

// MongoPipeline gets aggregation pipeline from a string
func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}

func toInt64(num interface{}) int64 {
	f := fmt.Sprintf("%v", num)
	x, _ := strconv.ParseFloat(f, 64)
	return int64(x)
}
