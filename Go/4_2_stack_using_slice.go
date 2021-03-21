// Implement stack using slice

package main

import "fmt"

func main() {
	slice := make([]int, 10, 20)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stack(slice[:])
}

func stack(s []int) {
	s = pop(s)
	s = push(s, 10)
	s = push(s, 11)
	s = push(s, 12)
	s = push(s, 13)
	s = push(s, 14)
	s = pop(s)
	s_top := top(s)
	fmt.Println(s_top)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func push(s []int, val int) []int {
	return append(s, val)
}

func top(s []int) int {
	return s[len(s)-1]
}
