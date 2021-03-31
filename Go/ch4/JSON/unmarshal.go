// JSON to Go DS

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"Released"`
	Color  bool `json: "color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casanova", Year: 2005, Color: true, Actors: []string{"Heath Ledger"}},
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {
	json := Marshal()
	Print()
	UnMarshal(json)
}

func UnMarshal(data []byte) {
	var titles []struct{ Title string }
	var actors []struct{ Actors []string }
	var year []struct{ Released int }
	fetch := []interface{}{titles, actors, year}
	if err := json.Unmarshal(data, &fetch); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(fetch...) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}

func Marshal() []byte {
	data, err := json.Marshal(movies) // No white spaces
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	return data
}

func Print() {
	data, err := json.MarshalIndent(movies, "", " ") // Human readable indentation
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
