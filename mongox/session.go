// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"context"
	"encoding/json"

	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Session -
type Session struct {
	collection *mongo.Collection
	ctx        context.Context
	filter     interface{}
	project    interface{}
	sort       interface{}
}

// Project sets sorting
func (s *Session) Project(project interface{}) *Session {
	s.project = project
	return s
}

// Sort sets sorting
func (s *Session) Sort(sort interface{}) *Session {
	s.sort = sort
	return s
}

// Decode returns all docs
func (s *Session) Decode(result interface{}) error {
	opts := options.Find()
	if s.sort != nil {
		opts.SetSort(s.sort)
	}
	if s.project != nil {
		opts.SetProjection(s.project)
	}
	cur, err := s.collection.Find(s.ctx, s.filter, opts)
	if err != nil {
		return err
	}

	var docs []bson.M
	for cur.Next(s.ctx) {
		var doc bson.M
		cur.Decode(&doc)
		docs = append(docs, doc)
	}
	b, _ := json.Marshal(docs)
	json.Unmarshal(b, result)
	return nil
}
