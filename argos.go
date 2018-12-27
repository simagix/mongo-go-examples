// Copyright 2018 Kuei-chun Chen. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"github.com/simagix/argos/examples"
	"github.com/simagix/keyhole/mdb"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("usage: argos uri collection pipeline")
		os.Exit(1)
	}

	var err error
	var client *mongo.Client
	var connStr connstring.ConnString
	if connStr, err = connstring.Parse(flag.Arg(0)); err != nil {
		panic(err)
	}

	if client, err = mdb.NewMongoClient(flag.Arg(0)); err != nil {
		panic(err)
	}

	var pipeline = []bson.D{}
	if len(flag.Args()) == 3 {
		if pipeline, err = mdb.GetAggregatePipeline(flag.Arg(2)); err != nil {
			panic(err)
		}
	}

	examples.ChangeStream(client, connStr.Database, flag.Arg(1), pipeline)
}
