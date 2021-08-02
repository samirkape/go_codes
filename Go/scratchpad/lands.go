package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/read", webBook)
	http.HandleFunc("/", welcome)
	http.ListenAndServe("", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my website")
	fmt.Fprintln(w, "Current time is: ", time.Now())
	fmt.Fprintln(w, "URL path is: ", r.URL.Path)
}

func webBook(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/yogesh/Downloads/Go_Web_Development.pdf")
}
