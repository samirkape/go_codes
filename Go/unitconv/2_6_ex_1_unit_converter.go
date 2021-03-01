// Unit conversion program currently supports temperature and length conversion

package main

import (
	"fmt"
	"gopl/unitconv"
)

func main() {

	select_unit()

}

func select_unit() {
	var choice int

	for {
		fmt.Println("Select Unit to Convert..")
		fmt.Println("1. Temperature")
		fmt.Println("2. Length")
		fmt.Println("-1. Exit")
		fmt.Print("Enter: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			select_unit_temp()
		case 2:
			fmt.Println("\nNot Implemented Yet..\n")
		default:
			return
		}

	}
}

func select_unit_temp() {

	var tempChoice int
	var tmp float64

	for {

		fmt.Println("Select Sub-Unit to Convert..")
		fmt.Println("Exit: Any other key")

		fmt.Println("1. Celsius to Fahrenheit")
		fmt.Println("2. Celsius to Kelvin")

		fmt.Println("3. Fahrenheit to Celsius")
		fmt.Println("4. Fahrenheit to Kelvin")

		fmt.Println("5. Kelvin to Celsius")
		fmt.Println("6. Kelvin to Fahrenheit")

		fmt.Print("Enter: ")
		fmt.Scan(&tempChoice)

		switch tempChoice {
		case 1:
			fmt.Println("Selected: 1. Celsius to Fahrenheit")
			fmt.Print("Enter Temperature in °C: ")
			fmt.Scan(&tmp)
			c := unitconv.Celsius(tmp)
			f := unitconv.TCtoF(c)
			fmt.Printf("\n%v °C are %v °F\n\n", c, f)

		case 2:
			fmt.Println("Selected: 2. Celsius to Kelvin")
			fmt.Print("Enter Temperature in °C: ")
			fmt.Scan(&tmp)
			c := unitconv.Celsius(tmp)
			k := unitconv.TCtoK(c)
			fmt.Printf("\n%v °C are %v °K\n\n", c, k)

		case 3:
			fmt.Println("Selected: 3. Fahrenheit to Celsius")
			fmt.Print("Enter Temperature in °F: ")
			fmt.Scan(&tmp)
			f := unitconv.Fahrenheit(tmp)
			c := unitconv.TFtoC(f)
			fmt.Printf("\n%v °F are %v °C\n\n", f, c)
		case 4:
			fmt.Println("Selected: 4. Fahrenheit to Kelvin")
			fmt.Print("Enter Temperature in °F: ")
			fmt.Scan(&tmp)
			f := unitconv.Fahrenheit(tmp)
			k := unitconv.TFtoK(f)
			fmt.Printf("\n%v °F are %v °K\n\n", f, k)
		case 5:
			fmt.Println("Selected: 5. Kelvin to Celsius")
			fmt.Print("Enter Temperature in °K\n\n: ")
			fmt.Scan(&tmp)
			k := unitconv.Kelvin(tmp)
			c := unitconv.TKtoC(k)
			fmt.Printf("\n%v °K are %v °C", k, c)
		case 6:
			fmt.Println("Selected: 6. Kelvin to Fahrenheit")
			fmt.Print("Enter Temperature in °K: ")
			fmt.Scan(&tmp)
			k := unitconv.Kelvin(tmp)
			f := unitconv.TKtoF(k)
			fmt.Printf("\n%v °K are %v °F\n\n", k, f)
		default:
			fmt.Println("Wrong Choice! Contine? 1/0")
			fmt.Print("Enter: ")
			var cont int
			fmt.Scan(&cont)
			if cont == 0 {
				return
			}
		}
	}

}
