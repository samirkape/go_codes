// Write an in-place function to eliminate adjacent duplicates in a []string slice

package main

import (
	"fmt"
)

func main() {
	str := "sameeeer kapee"
	slice := []rune(str)
	fmt.Println(string(remove_dup(slice[:])))
}

func remove_dup(str []rune) []rune {
	for i := 0; i+1 < len(str); i++ {
		if str[i] == str[i+1] {
			str = remove_adj(str, i)
			i--
		}
	}
	return str
}

func remove_adj(str []rune, index int) []rune {
	dest := index + 1
	src := index + 2
	copy(str[dest:], str[src:])
	return str[:len(str)-1]
}
