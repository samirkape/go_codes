//this program will be using Files (on CLA) to read text (io/ioutil) and report duplicates
// if no CLA is provided the program will go for reading text from stdin

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	store := make(map[string]int)
	args := os.Args[1:]

	for _, element_f := range args {
		data, err := ioutil.ReadFile(element_f)
		check_err(err)

		for _, word := range strings.Split(string(data), " ") {
			store[word]++
		}
	}
	for word, count := range store {
		if count > 1 {
			fmt.Printf("%s  ->  %d\n", word, count)
		}
	}

}

func check_err(err error) {
	if err != nil {
		fmt.Errorf("1-3-2: ", err)
	}
}
