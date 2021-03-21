// Rotate Right and Left by Reversing an array inplace using slices

package main

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	// rotate_left(arr[:], 1)
	rotate_right(arr[:], 2)
}

func rotate_left(rotate []int, hops int) {
	rev(rotate[:])
	rev(rotate[:len(rotate)-(hops)])
	rev(rotate[len(rotate)-hops:])
}

func rotate_right(rotate []int, hops int) {
	rev(rotate[:])
	rev(rotate[hops:])
}

func rev(rev []int) {
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
}
