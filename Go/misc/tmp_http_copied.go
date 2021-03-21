package main

import (
    "fmt"
    "net/http"
)

func hello_handler(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func main() {
	var i, j, k = true, x = 2.5, y, z = 1;  

    http.HandleFunc("/hello", hello_handler)
    http.ListenAndServe("localhost:8090", nil)
}