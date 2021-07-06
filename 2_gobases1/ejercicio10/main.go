package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La lista de empleados es: ", employees)
	fmt.Println("La edad de benjamin es: ", employees["Benjamin"])

	var totalAdults int = 0

	for key := range employees {
		if employees[key] > 21 {
			totalAdults++
		}
	}

	fmt.Println("La cantidad de empleados mayor a 21 años es: ", totalAdults)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println("La nueva lista de empleados es: ", employees)
}
