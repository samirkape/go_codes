package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandlerFunc("/product/{id")
	fmt.Println()
}
