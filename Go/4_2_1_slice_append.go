// The append function implementation

package main

func main() {
	// var slice []int = make([]int, 6, 6)
	slice := []int{1, 2, 3, 4, 5}
	slice = myappend(slice[:], 6)
}

func myappend(s []int, val int) []int {
	s_len := len(s) + 1
	s_cap := cap(s)
	var ns []int
	if s_len > s_cap {
		// allocate a new array with double size
		ns_cap := 2 * s_len
		ns = make([]int, s_len, ns_cap)
		copy(ns, s)
	} else {
		// this is how we can extend the slice
		ns = s[:s_len]
	}
	ns[s_len-1] = val
	return ns
}
