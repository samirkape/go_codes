/* The JSON-based web service of the Open Movie Database lets you search https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line. */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Poster string `json:"Poster"`
}

var URL = "https://omdbapi.com/"
var APIKey = "3e8be282"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
	response.Body.Close()
	return PosterURL.Poster
}

func RaiseFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	MovieName := os.Args[1]
	//MovieName := "casanova"
	url := BuildQuery(MovieName)
	PosterURL := FetchPosterURL(url)
	//EmptyFile, FileName := CreateFile(MovieName)
	FetchImg(w, PosterURL)
	//defer EmptyFile.Close()
	//fmt.Printf("\tDownload Successful: %s\n", FileName)
}

func CreateFile(MovieName string) (*os.File, string) {
	filename := strings.ReplaceAll(MovieName, " ", "+")
	filename += ".jpg"
	file, err := os.Create(filename)
	RaiseFatal(err)
	return file, filename
}

func FetchImg(w io.Writer, PosterURL string) {
	response, err := http.Get(PosterURL)
	if err != nil {
		response.Body.Close()
		log.Fatal(err)
	}
	// data, err := ioutil.ReadAll(response.Body)
	io.Copy(w, response.Body)
	if err != nil {
		response.Body.Close()
		log.Fatal(err)
	}
	response.Body.Close()
}

func BuildQuery(MovieName string) string {
	base, err := url.Parse(URL)
	RaiseFatal(err)
	params := url.Values{}
	params.Add("t", MovieName)
	params.Add("apikey", APIKey)
	base.RawQuery = params.Encode()
	fmt.Printf("Encoded URL is %q\n", base.String())
	return base.String()
}
