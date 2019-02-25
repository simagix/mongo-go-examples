// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/simagix/mongo-go-examples/examples"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://localhost/argos?replicaSet=replset&authSource=admin"

func seed() {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	examples.SeedCarsData(client, "argos")
}

func TestFindSort(t *testing.T) {
	var ctx = context.Background()
	client, err := Connect(ctx, uri)
	if err != nil {
		t.Fatal(err)
	}
	seed()
	c := client.Database("argos").Collection("cars")
	filter := bson.D{{Key: "color", Value: "Red"}}
	sort := bson.D{{Key: "style", Value: -1}}
	project := bson.D{{Key: "_id", Value: 0}, {Key: "style", Value: 1}, {Key: "brand", Value: 1}, {Key: "dealer", Value: 1}}
	var docs []bson.M
	if err = c.Find(filter).Project(project).Sort(sort).Limit(2).Decode(&docs); err != nil {
		t.Fatal(err)
	}
	t.Log(stringify(docs, "", "  "))
	var result []bson.M
	filter = bson.D{{Key: "style", Value: "Truck"}}
	project = bson.D{{Key: "_id", Value: 0}, {Key: "color", Value: 1}, {Key: "brand", Value: 1}, {Key: "dealer", Value: 1}}
	if err = c.Find(filter).Project(project).Skip(10).Limit(3).Decode(&result); err != nil {
		t.Fatal(err)
	}
	t.Log(stringify(result, "", "  "))
}

func stringify(doc interface{}, opts ...string) string {
	if len(opts) == 2 {
		b, _ := json.MarshalIndent(doc, opts[0], opts[1])
		return string(b)
	}
	b, _ := json.Marshal(doc)
	return string(b)
}
