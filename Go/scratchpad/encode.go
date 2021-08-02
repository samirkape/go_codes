package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	var Samir Person

	input := []byte(`{"Name":"Samir", "Age": 26}`)
	// inputJson := bytes.NewReader(input)

	json.NewDecoder(bytes.NewReader(input)).Decode(&Samir)
	fmt.Println(Samir)
}
