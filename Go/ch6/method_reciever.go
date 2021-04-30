package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance1(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance2(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Name string // you have to create a differenet type if you want to send a message to object

func (s Name) Distance3() Name {
	return s + Name(" Kape")
}

func main() {
	var n Name
	n = "Samir"
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("Distance1: ", Distance1(p, q))
	fmt.Println("Distance2: ", p.Distance2(q))
	fmt.Println("Name: ", n.Distance3())
}
