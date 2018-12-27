// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// ChangeStream prints oplogs in JSON format
func ChangeStream(client *mongo.Client, dbname string, collname string, pipeline []bson.D) {
	fmt.Println("Watching", dbname+"."+collname)
	var err error
	var coll = client.Database(dbname).Collection(collname)
	var ctx = context.Background()
	var cur mongo.Cursor
	fmt.Println("pipeline", pipeline)
	opts := options.ChangeStream()
	opts.SetFullDocument("updateLookup")
	if cur, err = coll.Watch(ctx, pipeline, opts); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	var b []byte
	var doc bson.M
	for cur.Next(ctx) {
		if err = cur.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		b, _ = json.MarshalIndent(doc, "", "  ")
		fmt.Println(string(b))
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
}
