// a program that counts the occurrences of each distinct Unicode code point in its input

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	count := make(map[rune]int)
	var utf8len [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		rn, nbytes, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if rn == unicode.ReplacementChar && nbytes == 1 {
			invalid++
			continue
		}
		count[rn]++
		utf8len[nbytes]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utf8len {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
