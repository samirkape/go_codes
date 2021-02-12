// this utility will fetch the URL specified on CL using ioutil.ReadAll
// refer >> man curl for more details

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	args := os.Args[1:]
	for _, url := range args {
		response, err := http.Get(url)
		check_err(err, "http GET error")
		response_body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		check_err(err, "reading error")
		fmt.Printf("%s ", response_body)
	}

}

func check_err(err error, estr string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %s %s ", estr, err)
		os.Exit(-2)
	}
}
