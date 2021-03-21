package main
//this code will print the name of this file with full path

import(
	"fmt"
	"os"
)

func main(){
	arg := os.Args[0]
	fmt.Println("__FILE__ = ", arg)
}