package main

import "fmt"

var record map[int]int

func filterMap(input ...int) {

	mk := input[0]
	mv := input[1]

	for k, v := range record {
		if k > mk && mv > k {
			delete(record, k)
			if v > record[mk] {
				record[mk] = v
			}
			filterMap(mk, v)
		}
	}
}

func filter(input [][]int) map[int]int {
	for _, i2 := range input {
		record[i2[0]] = i2[1]
	}

	for _, i2 := range input {
		filterMap(i2...)
	}

	return nil
}

func main() {

	//input := [][]int{{10, 12}, {1, 11}, {2, 13}, {22, 25}, {2, 13}} //, {20, 25}, {3, 8}, {6, 8}, {5, 6}}
	//input1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}           //, {20, 25}, {3, 8}, {6, 8}, {5, 6}}
	input1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {7, 21}} //, {20, 25}, {3, 8}, {6, 8}, {5, 6}}
	record = make(map[int]int)

	filter(input1)

	for k, v := range record {
		fmt.Printf("%d\t%d", k, v)
		fmt.Println()
	}
}
