package bsonu

import "encoding/json"

// M is a convenient alias for a map[string]interface{} map, useful for dealing with BSON in a native way. For instance:
type M map[string]interface{}

// BeautifyJSON returns beautified JSON
func BeautifyJSON(str string) (string, error) {
	bytes := []byte(str)
	var v = M{}
	json.Unmarshal(bytes, &v)
	b, err := json.MarshalIndent(v, "", "  ")
	return string(b), err
}
