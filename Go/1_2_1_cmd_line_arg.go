package main
//this code prints the command line arguments passed using range for loop 

import(
	"os"
	"fmt"
)

func main(){
	var s,sep = "","" 
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "  "
	}
	fmt.Println(s)
}
