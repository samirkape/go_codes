package main
// prints CLA using strings.Join()

import(
	"fmt"
	"os"
	"strings"
)

func main(){
	arg := os.Args[1:]
	fmt.Println(strings.Join(arg, " "))
}