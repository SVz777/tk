package json

import (
	"bytes"
	bjson "encoding/json"
)

type Number = bjson.Number

// Marshal ...
func Marshal(v interface{}) ([]byte, error) {
	return bjson.Marshal(v)
}

// Unmarshal 处理json float64精度丢失
func Unmarshal(data []byte, v interface{}) error {
	decoder := bjson.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err := decoder.Decode(&v)
	return err
}
