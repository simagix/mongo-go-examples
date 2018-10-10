package argos

import (
	"context"
	"fmt"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// example: argos "mongodb://localhost:27017/argos?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'
func TestPrintOplogs(t *testing.T) {
	uri := "mongodb://localhost:30097/argos?replicaSet=replset"
	pipelineStr := "[]"
	collname := "oplogs"
	connStr, e1 := connstring.Parse(uri)
	if e1 != nil {
		panic(e1)
	}
	client, e2 := mongo.Connect(context.Background(), connStr.String(), nil)
	if e2 != nil {
		panic(e2)
	}

	fmt.Println(connStr.Database, client)
	pipeline, e3 := bson.ParseExtJSONArray(pipelineStr)
	if e3 != nil {
		panic(e3)
	}

	go func() {
		PrintOpLogs(client, connStr.Database, collname, pipeline)
	}()
}
