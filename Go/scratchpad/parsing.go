// { "name":"samir","id":212}
package main

import (
	"bytes"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func main() {

	var NewSamir Person

	m := make(map[int]string)

	samir := Person{
		Name: "Samir",
		ID:   212,
	}

	inputb, err := json.Marshal(samir)

	if err != nil {
		return
	}

	inputr := bytes.NewReader(inputb)

	json.NewDecoder(inputr).Decode(&NewSamir)

}
