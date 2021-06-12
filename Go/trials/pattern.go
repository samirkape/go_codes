package main

import (
	"fmt"
	"strings"
)

func flip( toggle *int ) {
	if *toggle == 0 {
		*toggle = 1
	} else {
		*toggle = 0
	}
}

// set these values as a dimensions of pattern
const INNER = 21
const OUTER = 11

func main() {
	pattern := ""
	toggle := 0
	inner := INNER
	outer := OUTER
	spaces := ""

	for i := 0; i < outer; i++ {
		for j := 0; j < inner; j++ {
			if inner != INNER {
				spaces = strings.Repeat(" ", (INNER - inner) / 2.0)
			}
			flip(&toggle)
			pattern = pattern + fmt.Sprintf("%d", toggle)
		}
		inner -= 2
		fmt.Println(spaces + pattern)
		pattern = ""
		flip(&toggle)
	}
}
