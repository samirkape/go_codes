//Write a function that reports whether two strings are anagrams of each other,
//that is, the y contain the same letters in a different order

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	string1 := args[0]
	string2 := args[1]

	fmt.Println("Strings are Anagram", is_anagram(string1, string2))

}

func is_anagram(string1, string2 string) bool {

	if len(string1) != len(string2) {
		return false
	} else {
		for _, i := range string1 {
			if !strings.Contains(string2, string(i)) {
				return false
			}
		}
	}
	return true
}
