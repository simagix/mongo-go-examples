package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/argos/core"
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

	var pipelineStr string
	if len(flag.Args()) >= 3 {
		pipelineStr = flag.Arg(2)
	}
	var pipeline *bson.Array
	if pipelineStr == "" {
		pipelineStr = "[]"
	}
	pipeline, err = bson.ParseExtJSONArray(pipelineStr)
	if err != nil {
		panic(err)
	}
	argos.PrintOpLogs(client, connStr.Database, flag.Arg(1), pipeline)
}
