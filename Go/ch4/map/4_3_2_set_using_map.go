// create set using map

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	set_s := set()
	fmt.Println(set_s)
}

func set() []string {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	i := 0

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
		}
	}
	out := make([]string, len(seen))
	for key := range seen {
		out[i] = key
		i++
	}
	return out
}
