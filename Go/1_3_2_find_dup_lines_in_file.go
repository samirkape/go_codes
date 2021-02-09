package main

//this program will be using Files (on CLA) to read text and report duplicates
// if no CLA is provided the program will go for reading text from stdin

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	store := make(map[string]int)

	if len(os.Args) > 1 {
		cla_list := os.Args[1:]
		for _, element := range cla_list {
			fd, err := os.Open(element)
			check_err(err)
			scan_text(fd, store)
			fd.Close()
		}
	} else {
		scan_text(os.Stdin, store)
	}

	for key, element := range store {
		fmt.Printf("key -> %s \t count -> %d\n", key, element)
	}
}

func check_err(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v", err)
	}
}

func scan_text(fd *os.File, counts map[string]int) {

	input := bufio.NewScanner(fd)
	for input.Scan() {
		counts[input.Text()]++
	}

}
