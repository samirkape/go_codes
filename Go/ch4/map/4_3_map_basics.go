// Trying map data structure by counting ocurrances of the characters in a string

package main

import "fmt"

func main() {
	str := "sameer nitin kape"
	o_map := count_ocurrances(str)
	print_map(o_map)
}

func print_map(m map[string]int) {
	for key, value := range m {
		fmt.Printf("char  %s\t", key)
		fmt.Printf("count %d\t\n", value)
	}
}

func count_ocurrances(str string) map[string]int {
	m := make(map[string]int)
	for _, i := range str {
		m[string(i)]++
	}
	delete(m, " ")
	return m
}
