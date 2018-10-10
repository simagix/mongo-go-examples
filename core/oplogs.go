package argos

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/simagix/argos/bsonu"
)

// PrintOpLogs prints oplogs in JSON format
func PrintOpLogs(client *mongo.Client, dbname string, collname string, pipeline *bson.Array) {
	fmt.Println("Watching", dbname+"."+collname)
	db := client.Database(dbname)
	coll := db.Collection(collname)
	ctx := context.Background()

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

		str, _ := bsonu.BeautifyJSON(elem.ToExtJSON(true))
		fmt.Println(str)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
