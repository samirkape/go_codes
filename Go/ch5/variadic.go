// Variadic functions in Go

package main

import "fmt"

func f(...int) {}
func g([]int)  {}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", sum(s...))
	fmt.Printf("%d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("%d\n", slice_sum(s))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func slice_sum(vals []int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
