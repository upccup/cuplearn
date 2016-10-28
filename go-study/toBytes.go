package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	var jsonMap = []byte(`{
		"name": "upccup", "age": 12
		}`)
	fmt.Println(string(jsonMap))
	var out interface{}
	if err := json.Unmarshal(jsonMap, &out); err != nil {
		return
	}
	fmt.Printf("out %+v", out)

	byteList, err := GetBytes(out)
	if err != nil {
		fmt.Printf("bytes %s \n", err.Error())
		return
	}
	fmt.Printf("bytes %+v \n", byteList)
	j, err := json.Marshal(out)
	if err != nil {
		fmt.Printf("aaaa %s", err.Error())
		return
	}
	fmt.Println(string(j))
}
