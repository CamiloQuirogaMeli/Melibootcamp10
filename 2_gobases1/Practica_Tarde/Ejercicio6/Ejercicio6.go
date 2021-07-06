package main

import "fmt"

func main() {
	emp := "Benjamin"
	cant := 0
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("Item 1: La edad de %q es: %v \n", emp, employees[emp])

	for _, valor := range employees {
		if valor > 21 {
			cant++
		}
	}
	fmt.Printf("Item 2: Los empleados mayores de 21 años son %v\n", cant)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println("Item 3:")
	for key, valor := range employees {
		fmt.Println(key, valor)
	}
}
