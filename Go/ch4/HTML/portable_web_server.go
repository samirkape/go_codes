// Creating a web server to store http get result in local mongodb instance

package _ws

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func init() {

}

func handler(w http.ResponseWriter, r *http.Request) {

}

func CreateServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func BuildQuery(APIKey string, URL string, MovieName string) string {
	base, err := url.Parse(URL)
	RaiseFatal(err)
	params := url.Values{}
	params.Add("t", MovieName)
	params.Add("apikey", APIKey)
	base.RawQuery = params.Encode()
	fmt.Printf("Encoded URL is %q\n", base.String())
	return base.String()
}

func RaiseFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteMessage(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, s)
}
