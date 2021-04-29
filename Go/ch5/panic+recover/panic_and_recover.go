package main

import (
	"fmt"
)

// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
func soleTitle() {
	type bailout struct{}
	var b bailout
	defer func() {
		fmt.Println("I'm panicked but still came here")
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			fmt.Println("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	panic(b) // you can recover from panic by specifying reason
}

func main() {
	soleTitle()
}
