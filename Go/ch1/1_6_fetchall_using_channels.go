// this program will use channels to simultaneously fetch the https response of the sites provided as an input.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	args := os.Args[1:]

	ch := make(chan string)

	for _, _url := range args {
		url := add_prefix(_url)
		go fetch(url, ch) // start a goroutine
	}

	for range args {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2f elapsed: \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // passing err to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)

	response.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2f  %7d  %s", secs, nbytes, url)
}

func add_prefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}
