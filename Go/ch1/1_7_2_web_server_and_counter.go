// A server which displays http response of the given URL

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int
var url string

func main() {
	_url := os.Args[1]
	url = addPrefix(_url)
	http.HandleFunc("/sam", handler)   // each request calls handler, also increments hit count
	http.HandleFunc("/count", counter) // write to URL/count
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	response, err := http.Get(url)
	errCheck(err)
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Fprintf(w, "URL = %q\n", r.URL.Path)
	fmt.Fprintf(w, "Data = %s\n", data)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Hit count : %d\n", count)
	mu.Unlock()
}

func errCheck(err error) {
	if err != nil {
		fmt.Printf("Fetch: Error %d", err)
	}
}

func addPrefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}
