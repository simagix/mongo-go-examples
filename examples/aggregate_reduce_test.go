// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"testing"

	"github.com/simagix/keyhole/mdb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAggregateMatch(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()

	client = getMongoClient()
	defer client.Disconnect(ctx)
	seedFavoritesData(client, dbName)

	pipeline := `
	[{
		"$match": {
			"favoritesList.book": "Journey to the West"
		}
	}, {
		"$project": {
			"_id": 0,
			"favoritesList": 1
		}
	}, {
		"$unwind": {
			"path": "$favoritesList"
		}
	}, {
		"$match": {
			"favoritesList.book": "Journey to the West"
		}
	}]`
	collection = client.Database(dbName).Collection(collectionFavorites)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, mdb.MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	t.Log("total", total)
}

func TestAggregateRedact(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()

	client = getMongoClient()
	seedFavoritesData(client, dbName)

	pipeline := `[
	  {
	    "$match": {
	      "favoritesList.book": "Journey to the West"
	    }
	  }, {
	    "$project": {
	      "_id": 0,
	      "favoritesList": 1
	    }
	  }, {
	    "$redact": {
	      "$cond": {
	        "if": {
	          "$or": [
	            {
	              "$eq": ["$book", "Journey to the West"]
	            }, {
	              "$not": "$book"
	            }
	          ]
	        },
	        "then": "$$DESCEND",
	        "else": "$$PRUNE"
	      }
	    }
	  }, {
	    "$unwind": {
	      "path": "$favoritesList"
	    }
	  }
	]`
	collection = client.Database(dbName).Collection(collectionFavorites)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, mdb.MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	t.Log("total", total)
}

func TestAggregateFilter(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var cur *mongo.Cursor
	var ctx = context.Background()

	client = getMongoClient()
	seedFavoritesData(client, dbName)

	pipeline := `[
	  {
	    "$match": {
	      "favoritesList.book": "Journey to the West"
	    }
	  }, {
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
	      "_id": 0
	    }
	  }, {
	    "$unwind": {
	      "path": "$favoritesList"
	    }
	  }
	]`
	collection = client.Database(dbName).Collection(collectionFavorites)
	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, mdb.MongoPipeline(pipeline), opts); err != nil {
		t.Fatal(err)
	}
	defer cur.Close(ctx)
	total := 0
	for cur.Next(ctx) {
		total++
	}
	t.Log("total", total)
}
