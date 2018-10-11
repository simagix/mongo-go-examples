package bsonu

import (
	"encoding/json"

	"github.com/mongodb/mongo-go-driver/bson"
)

// NewDocument converts map[string]interface{} to bson.Document
func NewDocument(m M) (*bson.Document, error) {
	b, e1 := json.Marshal(m)
	if e1 != nil {
		return nil, e1
	}

	doc, err := bson.ParseExtJSONObject(string(b))
	return doc, err
}

// NewArray converts []map[string]interface{} to bson.Array
func NewArray(m []M) (*bson.Array, error) {
	b, e1 := json.Marshal(m)
	if e1 != nil {
		return nil, e1
	}

	array, err := bson.ParseExtJSONArray(string(b))
	return array, err
}
