// Copyright 2018 Kuei-chun Chen. All rights reserved.

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/network/connstring"
	"github.com/simagix/keyhole/mdb"
	"github.com/simagix/mongo-go-examples/examples"
)

func main() {
	var err error
	var client *mongo.Client
	var connStr connstring.ConnString
	collection := flag.String("collection", "", "collection to watch")
	pipe := flag.String("pipeline", "", "aggregation pipeline")
	caFile := flag.String("sslCAFile", "", "CA file")
	clientPEMFile := flag.String("sslPEMKeyFile", "", "client PEM file")

	flag.Parse()
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if connStr, err = connstring.Parse(flag.Arg(0)); err != nil {
		panic(err)
	}

	if client, err = mdb.NewMongoClient(flag.Arg(0), *caFile, *clientPEMFile); err != nil {
		panic(err)
	}

	var pipeline = []bson.D{}
	if *pipe != "" {
		pipeline = mdb.MongoPipeline(*pipe)
	}

	stream := examples.NewChangeStream()
	stream.SetCollection(*collection)
	stream.SetDatabase(connStr.Database)
	stream.SetPipeline(pipeline)
	stream.Watch(client, echo)
}

func echo(doc bson.M) {
	var b []byte
	b, _ = json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(b))
}
