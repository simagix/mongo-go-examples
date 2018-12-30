// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Collection contains mongo.Collection
type Collection struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find finds docs by given filter
func (c *Collection) Find(ctx context.Context, filter interface{}) *Session {
	return &Session{filter: filter, ctx: ctx, collection: c.collection}
}
