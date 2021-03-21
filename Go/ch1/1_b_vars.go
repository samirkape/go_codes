package main
// different types of variable declarations and recommended ones

import(
	"fmt"
)

func main(){
	
	s0 := "s0_printed"
	var s1 string
	var s2 = "s2_printed"
	var s3 string = "s3_printed"

	fmt.Println("var s0 := \" \" ->  ", s0)
	fmt.Println("var s1 string ->  ", s1)
	fmt.Println("var s2 = \" \" ->  ", s2)
	fmt.Println("var s3 string = \" \" ->  ", s3)
}
