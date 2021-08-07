package main

import "fmt"

func main() {
	Parent()
	fmt.Println("Grand-parent recovered")
}

func Parent() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	Child()
	fmt.Println("Parent recovered")
}

func Child() {
	panic("Hey")
}
