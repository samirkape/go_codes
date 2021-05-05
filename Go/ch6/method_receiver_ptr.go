// Creating a pointer receiver for converting termperature from Fahrenheit to Celsius

package main

import "fmt"

func main() {
	tempConv()
}

type Fahr []float64

func tempConv() {
	F := Fahr{103, 102, 100}
	fmt.Printf("%p\n", &F)
	F.Celsius()
	fmt.Println(F)
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
