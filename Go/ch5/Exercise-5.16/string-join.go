// Exercise 5.16: Write a variadic version of strings.Join.

package main

import "fmt"

type Name string

func main() {
	//slice := []string{"Samir", "Nitin", "Kape"}
	s := Join(",", "Samir", "Nitin", "Kape")
	fmt.Println(s)
}

func Join(del string, args ...string) (ret string) {
	for _, e := range args {
		ret += e
		ret += del
	}
	return ret[:len(ret)-1] // 0 to len(ret)-1 (exclusive)
}

// func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
