package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Edad de Benjamin: %d\n", employees["Benjamin"])

	var mayores21 int
	for _, age := range employees {
		if age > 21 {
			mayores21++
		}
	}
	fmt.Printf("Cantidad de empleados mayores a 21: %d\n", mayores21)

	employees["Federico"] = 25
	delete(employees, "Pedro")
}
