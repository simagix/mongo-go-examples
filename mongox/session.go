// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongox

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Session -
type Session struct {
	collection *mongo.Collection
	filter     interface{}
	limit      int64
	project    interface{}
	skip       int64
	sort       interface{}
}

// Limit sets sorting
func (s *Session) Limit(limit int64) *Session {
	s.limit = limit
	return s
}

// Project sets sorting
func (s *Session) Project(project interface{}) *Session {
	s.project = project
	return s
}

// Skip sets sorting
func (s *Session) Skip(skip int64) *Session {
	s.skip = skip
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
	if s.limit > 0 {
		opts.SetLimit(s.limit)
	}
	if s.project != nil {
		opts.SetProjection(s.project)
	}
	if s.skip > 0 {
		opts.SetSkip(s.skip)
	}
	if s.sort != nil {
		opts.SetSort(s.sort)
	}
	cur, err := s.collection.Find(nil, s.filter, opts)
	if err != nil {
		return err
	}

	var docs []bson.M
	for cur.Next(nil) {
		var doc bson.M
		cur.Decode(&doc)
		docs = append(docs, doc)
	}
	b, _ := json.Marshal(docs)
	json.Unmarshal(b, result)
	return nil
}
