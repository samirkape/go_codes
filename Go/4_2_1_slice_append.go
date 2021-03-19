// The append function implementation

package main

func main() {
	var slice []int = make([]int, 5, 6)
	//slice := []int{1, 2, 3, 4, 5}
	slice = myappend(slice[:], 6)

	v_slice := []int{1, 2, 3, 4, 5}
	v_slice = myappend_variadic(v_slice, v_slice...)
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

func myappend_variadic(s []int, val ...int) []int {
	s_len := len(s) + len(val)
	s_cap := cap(s)
	var ns []int
	if s_len > s_cap {
		ns_len := 2 * len(s)
		ns_cap := 0
		if s_cap < ns_len {
			ns_cap = ns_len
		}
		ns = make([]int, ns_len, ns_cap)
		copy(ns, s)
	} else {
		ns = s[:s_len]
	}
	copy(ns[len(s):], val)
	return ns
}
