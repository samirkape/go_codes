package main
// this code will fetch string lines from stdin
// will be using map to count character occurances in the string 

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
