package argos

import (
	"context"
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// example: argos "mongodb://localhost:27017/argos?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'
func TestPrintOplogs(t *testing.T) {
	uri := "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	pipelineStr := "[]"
	collname := "oplogs"
	connStr, e1 := connstring.Parse(uri)
	if e1 != nil {
		t.Fatal(e1)
	}
	client, e2 := mongo.Connect(context.Background(), connStr.String(), nil)
	if e2 != nil {
		t.Fatal(e2)
	}

	t.Log(connStr.Database)
	pipeline, e3 := bson.ParseExtJSONArray(pipelineStr)
	if e3 != nil {
		t.Fatal(e3)
	}

	go func() {
		PrintOpLogs(client, connStr.Database, collname, pipeline)
	}()
}
