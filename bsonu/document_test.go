package bsonu

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewDocument(t *testing.T) {
	m := M{"name": "MongoDB", "release": "4.0.3"}
	a, e1 := json.Marshal(m)
	if e1 != nil {
		t.Fatal(e1)
	}

	doc, e2 := NewDocument(m)
	if e2 != nil {
		t.Fatal(e2)
	}

	str := doc.ToExtJSON(true)
	obj := M{}
	json.Unmarshal([]byte(str), &obj)
	b, _ := json.Marshal(obj)

	if bytes.Compare(a, b) != 0 {
		t.Log(string(a))
		t.Log(string(b))
		t.Fail()
	}
}
