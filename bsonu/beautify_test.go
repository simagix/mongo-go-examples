package bsonu

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBeautifyJSON(t *testing.T) {
	str := "{\"name\": \"MongoDB\", \"release\": \"4.0.3\", \"scores\": [100, 95, 88]}"
	obj := M{}
	json.Unmarshal([]byte(str), &obj)
	a, _ := json.Marshal(obj)
	t.Log(str)
	bstr, _ := BeautifyJSON(str)
	bobj := M{}
	t.Log(bstr)
	json.Unmarshal([]byte(bstr), &bobj)
	b, _ := json.Marshal(bobj)

	if bytes.Compare(a, b) != 0 {
		t.Fail()
	}
}
