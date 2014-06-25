package main

import (
	"io/ioutil"
	"encoding/json"
)

type TestJson struct {
	Hoge string `json:"hoge"`
}

func Hoge() TestJson {
	b, err := ioutil.ReadFile("/workspace/test.json")

	if err != nil {
		panic(err)
	}
	var testJson TestJson
	json.Unmarshal(b, &testJson)

	return testJson
}
