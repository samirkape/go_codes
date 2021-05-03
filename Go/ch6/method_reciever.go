// Using methods

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 { // Notice Distance( ) with multiple definitions.
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Name string // You have to create a different type if you want to send a message to an object with primitive underlying type

func (s Name) Distance() Name {
	return s + Name(" Kape")
}

func main() {
	var n Name
	n = "Samir"
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("Distance1: ", Distance(p, q))
	fmt.Println("Distance2: ", p.Distance(q))
	fmt.Println("Name: ", n.Distance())
}
