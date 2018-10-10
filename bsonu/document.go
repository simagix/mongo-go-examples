package bsonu

import (
	"encoding/json"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

// NewDocument converts map[string]interface{} to bson.Document
func NewDocument(m M) (*bson.Document, error) {
	b, e1 := json.Marshal(m)
	fmt.Println("bytes", string(b))
	if e1 != nil {
		return nil, e1
	}

	doc, err := bson.ParseExtJSONObject(string(b))
	return doc, err
}
