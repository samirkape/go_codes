/* The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line. */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Movie struct {
	Poster string `json:"Poster"`
}

var URL = "https://omdbapi.com/"
var APIKey = "3e8be282"

func main() {
	// MovieName := os.Args[1]
	MovieName := "casanova"
	// col := InitDB()
	//DeleteCollection(col)
	url := BuildQuery(MovieName)
	PosterURL := FetchPosterURL(url)
	fmt.Printf("\tURL\t\t\ttranscript\n")
	fmt.Println(PosterURL)
}

func FetchPosterURL(rawurl string) string {
	response, err := http.Get(rawurl)
	var PosterURL Movie
	if err != nil && response.StatusCode > 300 {
		response.Body.Close()
		log.Fatal(err)
	}
	if err := json.NewDecoder(response.Body).Decode(&PosterURL); err != nil {
		response.Body.Close()
		log.Fatal(err)
	}
	return PosterURL.Poster
}

func RaiseFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func BuildQuery(MovieName string) string {
	base, err := url.Parse(URL)
	RaiseFatal(err)
	// Path params
	//base.Path += "this will get automatically encoded"
	// Query params
	params := url.Values{}
	params.Add("t", MovieName)
	params.Add("apikey", APIKey)
	base.RawQuery = params.Encode()
	fmt.Printf("Encoded URL is %q\n", base.String())
	return base.String()
}
