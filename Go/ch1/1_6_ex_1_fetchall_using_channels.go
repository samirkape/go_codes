// this program will fetch the URLs concurrently using goroutine and channels
// I will be using file to write website response body

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	args := os.Args[1:]
	ch := make(chan string)
	for _, _url := range args {
		url := add_prefix(_url)
		go fetchall(ch, url)
	}
	for i := 0; i < len(args); i++ {
		name := strconv.Itoa(i) + "_url.txt"
		file, err := os.Create(name)
		if err != nil {
			fmt.Printf("Fetch: Error in file open %s", err)
		}
		data := <-ch
		fmt.Printf("%s\n", data)
		fmt.Fprintf(file, "%s\n", data)
		file.Close()
	}

}

func fetchall(ch chan<- string, url string) {

	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Fetch: Error in http get %s", err)
	}
	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	ch <- fmt.Sprintf("%.2f\t%d\t%s", time.Since(start).Seconds(), nbytes, url)
}

func add_prefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}
