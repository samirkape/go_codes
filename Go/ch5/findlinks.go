// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// filename := fmt.Sprintf("%s", os.Args[1:])
	filename := fmt.Sprintf("%s", "./golang.org")
	doc, err := html.Parse(Open(filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func OpenFile(filename string) io.Reader {
	OpenFile, _ := ioutil.ReadFile(filename)
	return bytes.NewReader(OpenFile)
}

func Open(filename string) (file io.Reader) {
	file, _ = os.Open(filename)
	return file
}
