// Creating a web server to store http get result in local mongodb instance

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"bookfetch"
)

var APIKey = bookfetch.APIKey
var URL = bookfetch.URL

func main() {
	CreateServer()
}

func handler(w http.ResponseWriter, r *http.Request) {
	var Book bookfetch.Books
	BookName, ok := r.URL.Query()["name"]
	if !ok || BookName == nil {
		WriteMessage(w, "Invalid Query")
		return
	}
	_url := BuildQuery(fmt.Sprintf("%s", BookName))
	response, err := http.Get(_url)
	if err != nil || response.StatusCode > 300 {
		WriteMessage(w, "Book Not found")
		return
	}
	if err := json.NewDecoder(response.Body).Decode(&Book); err != nil {
		response.Body.Close()
		return
	}

	EditBookMeta(&Book)

	var BookMeta = bookfetch.BookLinks
	if err := BookMeta.Execute(w, Book); err != nil {
		log.Fatal(err)
	}
}

func EditBookMeta(Book *bookfetch.Books) {
	Book.TotalItems = len(Book.Items)
	for _, i := range Book.Items {
		if len(i.VolumeInfo.ImageLinks.Thumbnail) == 0 {
			i.VolumeInfo.ImageLinks.Thumbnail = ""
		}
	}
}

func CreateServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func BuildQuery(Name string) string {
	base, err := url.Parse(URL)
	RaiseFatal(err)
	params := url.Values{}
	params.Add("key", APIKey)
	params.Add("q", Name)
	base.RawQuery = params.Encode()
	fmt.Printf("Encoded URL is %q\n", base.String())
	return base.String()
}

func RaiseFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteMessage(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, s)
}
