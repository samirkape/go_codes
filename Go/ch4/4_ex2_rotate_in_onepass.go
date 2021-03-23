// Rotate Right and Left by Using Copy

package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	slice_l := rotate_left(arr, 3)
	fmt.Println(slice_l)
	slice_r := rotate_right(arr, 2)
	fmt.Println(slice_r)
}

func rotate_left(a [6]int, pos int) []int {
	s := a[:]
	var tmp []int = make([]int, pos)
	copy(tmp, s[:pos])
	copy(s[:], s[pos:])
	copy(s[len(s)-pos:], tmp)
	return s
}

func rotate_right(a [6]int, pos int) []int {
	s := a[:]
	var tmp []int = make([]int, pos)
	copy(tmp, s[len(s)-pos:])
	copy(s[pos:], s[:])
	copy(s[:], tmp)
	return s
}
