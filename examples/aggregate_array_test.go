// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/simagix/keyhole/mdb"
)

/*
 * find people who live in London and like the book "Journey to the West"
 * only displays matched.
 */
func TestAggregateArray(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur mongo.Cursor
	var ctx = context.Background()
	var doc bson.M

	client = getMongoClient()
	seedFavoritesData(client, dbName)

	pipeline := `[
		{
			"$match": {
			  "favoritesList": {
			    "$elemMatch": {
			      "city": "London",
			      "book": "Journey to the West"
			    }
			  }
		}}, {
			"$project": {
		     "favoritesList": {
		        "$filter": {
		          "input": "$favoritesList",
		          "as": "favorite",
		          "cond": {
		            "$eq": ["$$favorite.book", "Journey to the West"]
		          }
		        }
		      },
		      "_id": 0,
		      "email": 1
		}}, {
			"$unwind": {
        "path": "$favoritesList"
    }}]`
	collection = client.Database(dbName).Collection(collectionFavorites)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, mdb.MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		cur.Decode(&doc)
		t.Log(doc["email"], "likes movie", "'", doc["favoritesList"].(bson.M)["movie"], "' too.")
		total++
	}
	t.Log("total", total)
}
