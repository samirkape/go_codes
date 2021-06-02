package main

import "fmt"

func intSeq() func() int {
	i := 0              // why instance of this variable is not getting newly created at each call? why it is preserving its old value.
	return func() int { // why a call on line 17 is starting directly from this line and not ln 6?
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // why it had destroyed the previous value of i in the intSeq?
	fmt.Println(newInts())
	fmt.Println(nextInt())
}
