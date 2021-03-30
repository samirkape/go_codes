// Accessing struct from differnt package

package main

import (
	"employee"
	"fmt"
)

func main() {
	// var samir = new(employee.Personal)
	var samir = employee.Personal{Name: "Sam"}
	fmt.Println(samir)
	set_personal_fields(samir)
	fmt.Println(samir)
}

func set_personal_fields(samir employee.Personal) {
	samir.Name = "Samir"
	samir.Father = "Nitin"
	samir.Sirname = "Kape"
	samir.Address = "Z.P School, Kokamthan"
}
