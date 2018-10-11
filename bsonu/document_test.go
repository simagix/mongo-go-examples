package bsonu

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func TestNewDocument(t *testing.T) {
	m := M{"name": "MongoDB", "release": "4.0.3"}
	a, e1 := json.Marshal(m)
	if e1 != nil {
		t.Fatal(e1)
	}

	doc, e2 := NewDocument(m)
	if e2 != nil {
		t.Fatal(e2)
	}

	str := doc.ToExtJSON(true)
	obj := M{}
	json.Unmarshal([]byte(str), &obj)
	b, _ := json.Marshal(obj)

	if bytes.Compare(a, b) != 0 {
		t.Log(string(a))
		t.Log(string(b))
		t.Fail()
	}
}

func TestNewArray(t *testing.T) {
	ctx := context.Background()
	uri := "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	client, err := mongo.Connect(ctx, uri)
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("argos")
	coll := db.Collection("oplogs")
	pipeline := []M{}
	pipeline = append(pipeline, M{"$match": M{"_id": "30097"}})
	pipeline = append(pipeline, M{"$project": M{"_id": 0}})
	array, e2 := NewArray(pipeline)
	if e2 != nil {
		t.Fatal(e2)
	}
	cur, e3 := coll.Aggregate(ctx, array)
	if e3 != nil {
		t.Fatal(e3)
	}
	for cur.Next(nil) {
		item := M{}
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		t.Log(item)
	}
	cur.Close(ctx)
}
