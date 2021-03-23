// Rotate Right and Left by Reversing an array inplace using slices

package main

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	reverse_arr(arr)
	reverse_ptr(&arr)
}

func reverse_ptr(arr *[6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverse_arr(arr [6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
