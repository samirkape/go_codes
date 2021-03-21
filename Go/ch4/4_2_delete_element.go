// Slice: Delete element from given position

package main

func main() {
	pos := 1
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = del(slice[:], pos)
}

func del(s []int, pos int) []int {
	if pos < 1 {
		return s
	}
	pos -= 1
	copy(s[pos:], s[pos+1:])
	return s[:len(s)-1]
}
