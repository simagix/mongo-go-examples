// Copyright 2018 Kuei-chun Chen. All rights reserved.

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"github.com/simagix/keyhole/mdb"
	"github.com/simagix/mongo-go-examples/examples"
)

func main() {
	var err error
	var client *mongo.Client
	var connStr connstring.ConnString
	pipe := flag.String("pipeline", "", "aggregation pipeline")
	collection := flag.String("collection", "", "collection to watch")

	flag.Parse()
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if connStr, err = connstring.Parse(flag.Arg(0)); err != nil {
		panic(err)
	}

	if client, err = mdb.NewMongoClient(flag.Arg(0)); err != nil {
		panic(err)
	}

	var pipeline = []bson.D{}
	if *pipe != "" {
		pipeline = mdb.GetAggregatePipeline(*pipe)
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
