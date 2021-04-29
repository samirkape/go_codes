// Example panic and recover behaviour
package main

import "fmt"

var arr [4]int

func main() {
	two()
	fmt.Println("Exited Normally from one()")
}

func two() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in two() ", r)
		}
	}()
	fmt.Println("Calling three...")
	three(0)
	fmt.Println("Exited Normally from three()")
}

func three(i int) {
	if i > 3 {
		fmt.Println("Panicking...")
		arr[i] = 5000
		//panic(fmt.Sprintf("\tvalue at the time of panic: %v", i))
	}
	fmt.Println("three: ", i)
	defer fmt.Println("defer: three: ", i)
	three(i + 1)
}
