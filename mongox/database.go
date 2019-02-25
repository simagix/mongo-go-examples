// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import "go.mongodb.org/mongo-driver/mongo"

// Database contains mongo.Database
type Database struct {
	database *mongo.Database
}

// Collection returns database
func (d *Database) Collection(collection string) *Collection {
	return &Collection{collection: d.database.Collection(collection)}
}
