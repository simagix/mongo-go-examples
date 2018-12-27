// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/network/connstring"
)

// example: argos "mongodb://localhost:27017/argos?replicaSet=replset" students '[{"$match": {"operationType": "update"}}]'
func TestPrintOplogs(t *testing.T) {
	uri := "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	in := []byte(`[]`)
	var raw bson.D
	json.Unmarshal(in, &raw)
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
	pipeline := mongo.Pipeline{}

	go func() {
		ChangeStream(client, connStr.Database, collname, pipeline)
	}()
}
