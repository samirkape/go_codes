package main

import "fmt"

func main() {
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p) // "2"
	x := mult(*p, *p)
	fmt.Println(x) // "2"

}

func mult(left int, new int) int {
	ptr := new(int)
	return left * new
}
