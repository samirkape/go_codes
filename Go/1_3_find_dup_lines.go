package main
// this code will find duplicate lines from stdin
// will be using bufio and map and make

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\n%d\t%s\n", n, line)
		}
	}
}