package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/core/connstring"
	argos "github.com/simagix/argos/core"
)

// example: argos "mongodb://localhost:27017/prudential?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'
func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("usage: argos uri collection")
		os.Exit(1)
	}

	connStr, err := connstring.Parse(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	var pipeline string
	if len(flag.Args()) >= 3 {
		pipeline = flag.Arg(2)
		fmt.Println("pipeline:", pipeline)
	}
	argos.PrintOpLogs(connStr, flag.Arg(1), pipeline)
}
