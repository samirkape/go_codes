package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Name"`
	ID   int    `json:"ID"`
}

func Mymarshal(p Person) string {
	bj, _ := json.Marshal(p)
	return string(bj)
}

func MyUnmarshal(jsonSInput string) Person {
	var p Person
	json.Unmarshal([]byte(jsonSInput), &p)
	return p
}

func MyUnmarshalGen(jsonSInput string) interface{} {
	var v interface{}
	json.Unmarshal([]byte(jsonSInput), &v)
	return v
}

func MyDecode(jsonSInput string) Person {
	var p Person
	jsonBinput := bytes.NewReader([]byte(jsonSInput))
	json.NewDecoder(jsonBinput).Decode(&p)
	return p
}

func main() {
	tmp := Person{
		ID: 1920,
	}

	json := Mymarshal(tmp)
	sam := MyUnmarshalGen(json)
	samD := MyDecode(json)

	for k, v := range sam.(map[string]interface{}) {
		fmt.Println(k, v)
	}

	fmt.Println(json)
	fmt.Println(sam)
	fmt.Println(samD)
}
