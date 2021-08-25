package main

import "fmt"

var Glob []int

func fib(num int) []int {

	curr, prev := 1, 1
	var out []int

	out = append(out, prev)
	out = append(out, curr)

	for i := 1; i <= num; i++ {
		out = append(out, curr+prev) // 2
		tmp := prev                  // 1
		prev = curr                  // 1
		curr = curr + tmp            // 2
	}

	return out
}

func main() {
	out := fib(10)
	Glob = out
	fmt.Println(Glob)
}
