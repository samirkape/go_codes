// Compare two maps

package main

import "fmt"

func main() {
	str := "sameer nitin kape"
	str1 := "kape nitin deoram"

	sk_map := count_ocurrances(str)
	nk_map := count_ocurrances(str1)

	fmt.Println(areEqual(sk_map, nk_map))
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

func areEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for m1k, m1v := range m1 {
		if m2v, ok := m2[m1k]; !ok || (m1v != m2v) {
			return false
		}
	}
	return true
}
