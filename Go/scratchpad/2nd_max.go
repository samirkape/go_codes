package main

import (
	"fmt"
	"math"
)

func secondMax(arr []int) int {
	max := math.MinInt64
	prevMax := math.MinInt64
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			prevMax = max
			max = arr[i]
		} else if arr[i] > prevMax {
			prevMax = arr[i]
		}
	}
	return prevMax
}

func main() {
	arr := []int{12, 5, 1, 2, 4, 3, 6, 8, 9, 7, 9, 9, 9}
	max := secondMax(arr)
	fmt.Printf("Second max is: %d\n", max)
}
