// Fetch downloads the contents of URL and returns the length and name of the local file

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Fetch: Error: Unable to get requested URL")
		return " ", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	file, err := os.Create(local)
	filename = local
	if err != nil {
		log.Printf("Fetch: Error: Unable to write file")
		return " ", 0, err
	}
	n, err = io.Copy(file, resp.Body)
	defer file.Close()
	if err != nil {
		log.Printf("Fetch: Error: Copy to file failed")
		return " ", 0, err
	}
	return
}

func main() {
	// args := os.Args[1:]
	args := "http://google.in"
	filename, length, err := fetch(args)
	if err != nil {
		log.Fatal("Fetch: Error: ", err)
	}
	log.Printf("Filename: %s", filename)
	log.Printf("Length: %d", length)
}
