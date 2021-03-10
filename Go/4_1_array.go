// Demonstrating Array usage
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EURO
	GBP
	RMB
)

func main() {

	fmt.Println("Declare an array")
	fmt.Println("var iarr [5]int\n")
	fmt.Println("find len in number of elements")
	fmt.Println("len(arr)")

	fmt.Println("\nInitialize array literal")
	fmt.Println("var q [3]int = [3]int{1, 2, 3}")
	fmt.Println("\nFind type of q")
	fmt.Println("%T")

	fmt.Println("\nq -- Re-assign declared array with the same size")
	fmt.Println("q = [3]int{4, 5, 6}\n")

	symbols := [...]string{USD: "$", EURO: "9", GBP: "£", RMB: "¥"}
	fmt.Println(GBP, symbols[GBP])
	fmt.Printf("%d\n", len(symbols))
}
