// this utility will fetch the URL specified on CL with string.HasPrefix
// refer >> man curl for more details

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	for _, _url := range args {
		url := add_prefix(_url)
		response, err := http.Get(url)
		check_err(err, "http GET error")
		io.Copy(os.Stdout, response.Body)
		response.Body.Close()
	}

}

func add_prefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		fmt.Println("http:// not found. adding...")
		return "http://" + url
	}
}

func check_err(err error, estr string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %s %s ", estr, err)
		os.Exit(-2)
	}
}
