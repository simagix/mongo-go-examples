package argos

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// PrintOpLogs prints oplogs in JSON format
func PrintOpLogs(connStr connstring.ConnString, collname string, pipelineStr string) {
	client, err := mongo.Connect(context.Background(), connStr.String(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Watching", connStr.Database+"."+collname)
	db := client.Database(connStr.Database)
	coll := db.Collection(collname)
	ctx := context.Background()
	var pipeline interface{}
	if pipelineStr != "" {
		pipeline, _ = bson.ParseExtJSONArray(pipelineStr)
	}
	fmt.Println("pipeline", pipeline)
	cur, err := coll.Watch(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		elem := bson.NewDocument()

		if err := cur.Decode(elem); err != nil {
			log.Fatal(err)
		}

		bytes := []byte(elem.ToExtJSON(true))
		var v = M{}
		json.Unmarshal(bytes, &v)
		bytes, _ = json.MarshalIndent(v, "", "  ")
		fmt.Println(string(bytes))
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
