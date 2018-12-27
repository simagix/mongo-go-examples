// Copyright 2018 Kuei-chun Chen. All rights reserved.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
	"github.com/simagix/argos/examples"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("usage: argos uri collection pipeline")
		os.Exit(1)
	}

	connStr, err := connstring.Parse(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	client, err := mongo.Connect(context.Background(), connStr.String(), nil)
	if err != nil {
		panic(err)
	}

	pipeline := mongo.Pipeline{}
	if err != nil {
		panic(err)
	}
	examples.ChangeStream(client, connStr.Database, flag.Arg(1), pipeline)
}
