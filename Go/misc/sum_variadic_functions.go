// Trying variadic function for calculating sum

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	//sum_it := []int{1, 2, 3, 4, 5}
	s := sum(stoi(args)...)
	fmt.Println(s)
}

func stoi(s []string) []int {
	iarr := make([]int, len(s))
	for i, e := range s {
		iarr[i], _ = strconv.Atoi(e)
	}
	return iarr
}

func sum(val ...int) int {
	s := 0
	for i := 0; i < len(val); i++ {
		s += val[i]
	}
	return s
}
