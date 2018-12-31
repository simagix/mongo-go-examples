// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/mongo-go-examples/examples"
)

var uri = "mongodb://localhost/argos?replicaSet=replset"

func seed() {
	client, _ := mongo.Connect(context.Background(), uri)
	examples.SeedCarsData(client, "argos")
}

func TestFindSort(t *testing.T) {
	var ctx = context.Background()
	client, err := Connect(ctx, uri)
	if err != nil {
		t.Fatal(err)
	}
	seed()
	filter := bson.D{{Key: "color", Value: "Red"}}
	sort := bson.D{{Key: "style", Value: -1}}
	project := bson.D{{Key: "_id", Value: 0}, {Key: "style", Value: 1}, {Key: "brand", Value: 1}, {Key: "dealer", Value: 1}}
	var docs []bson.M
	err = client.Database("argos").Collection("cars").Find(filter).Project(project).Sort(sort).Decode(&docs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(stringify(docs, "", "  "))
}

func stringify(doc interface{}, opts ...string) string {
	if len(opts) == 2 {
		b, _ := json.MarshalIndent(doc, opts[0], opts[1])
		return string(b)
	}
	b, _ := json.Marshal(doc)
	return string(b)
}
