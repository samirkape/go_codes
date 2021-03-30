// How to initialize nested structs

package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	method1()
	method2()
}

func method1() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
}

func method2() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}
