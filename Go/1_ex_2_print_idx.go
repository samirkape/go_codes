package main

import(
	"fmt"
	"os"
)

func main(){
	arg := os.Args[1:]
	for idx, args := range arg{
		fmt.Println(idx,"-> ", args)
	} 
}