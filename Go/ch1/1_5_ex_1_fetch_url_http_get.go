// this utility will fetch the URL specified on CL with io.Copy() and os.Stdout
// refer >> man curl for more details

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	args := os.Args[1:]
	for _, url := range args {
		response, err := http.Get(url)
		check_err(err, "http GET error")
		//response_body, err := ioutil.ReadAll(response.Body)
		io.Copy(os.Stdout, response.Body)
		//fmt.Printf("%s ", response_body)
		response.Body.Close()
	}

}

func check_err(err error, estr string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %s %s ", estr, err)
		os.Exit(-2)
	}
}
