// Rough codes separated by the functions

package main

import "fmt"

func main() {
	hw := "こんにちは世界"
	fmt.Println(hw)

	slice_append_variac()
}

func slice_append_variac() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, s...) // 3 dots are neccesary if you want to append a slice into the slice
}
