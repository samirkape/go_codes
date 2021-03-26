// Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wct := make(map[string]int)

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wct[scanner.Text()]++
	}

	for k, v := range wct {
		fmt.Printf("%s\t\t%d\n", k, v)
	}
}
