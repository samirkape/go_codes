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
	ocur := make(map[byte]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		text := input.Text()
		counts[text]++
		for i := 0; i < len(text); i++{
			ocur[text[i]]++
		}
	}	

	for line, n := range ocur {
		fmt.Printf("char -> %c\tcount -> %d\n",line, n)
	}

}