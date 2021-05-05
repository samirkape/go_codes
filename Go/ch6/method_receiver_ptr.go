// Creating a pointer receiver for converting temperature from Fahrenheit to Celsius

package main

import "fmt"

func main() {
	tempConv()
}

type Fahr []float64
type Cels []float64

func tempConv() {
	F := Fahr{103, 102, 100}
	fmt.Printf("%p\n", &F)
	C := F.CelsiusR() // Getting C in the return, preserving F
	F.Celsius()       // Modifying F
	fmt.Println(F)
	fmt.Println(C)
	fmt.Printf("%T\n", F)
}

func (F *Fahr) Celsius() {
	fmt.Printf("%p\n", F)
	F1 := *F
	for i := 0; i < len(*F); i++ { // To get the length of pointer variable you need to dereference it
		F1[i] = (F1[i] - 32) * 9 / 5
	}
	F = &F1
}

func (F Fahr) CelsiusR() (C Cels) { // Named return too need allocation before using
	C = make(Cels, len(F))
	fmt.Printf("%p\n", F)
	for i, f := range F {
		C[i] = (f - 32) * 9 / 5
	}
	return
}
