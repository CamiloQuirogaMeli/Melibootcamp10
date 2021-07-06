package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println("Lista de empleados:")
	for name, age := range employees {
		fmt.Println(" *", name, age)
	}
	fmt.Println()
	fmt.Println("Empleados mayores a 21 años:")
	for name, age := range employees {
		if age >= 21 {
			fmt.Println(" *", name, age)
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println()
	fmt.Println("Nueva lista de empleados:")
	for name, age := range employees {
		fmt.Println(" *", name, age)
	}
}
