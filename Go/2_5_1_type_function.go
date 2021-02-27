// This program adds suffix °C to temp by suing var.String()

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero = -273.15
	Freezing     = 0
	Boiling      = 100
)

func CtoF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 9 / 5) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	c := FtoC(28)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
}
