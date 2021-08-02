package main

import "fmt"

func main() {
	a()
	fmt.Println("main function")
}

func a() {
	defer func() {
		fmt.Println("function a")
	}()
	b()
	return
}

func b() {
	panic(1)
	fmt.Println("function b")
	return
}
