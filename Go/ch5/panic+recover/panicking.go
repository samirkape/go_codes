package main

import "fmt"

func main() {
	defer func() { fmt.Println("Freeing Grand Father") }()
	child()
}

func child() {
	defer func() { fmt.Println("Freeing Father") }()
	grandchild()
}

func grandchild() {
	defer func() { fmt.Println("Freeing Me") }()
	panic("Panicking...")
	fmt.Println("Panic stops normal flow of execution, hence this line will never execute")
}
