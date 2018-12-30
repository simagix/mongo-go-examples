// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Client contains mongo.Client
type Client struct {
	client *mongo.Client
}

// Connect creates a new Client and then initializes it using the Connect method.
func Connect(ctx context.Context, uri string) (*Client, error) {
	var err error
	var client *mongo.Client
	client, err = mongo.Connect(ctx, uri)
	return &Client{client: client}, err
}

// Database returns database
func (c *Client) Database(db string) *Database {
	return &Database{database: c.client.Database(db)}
}
