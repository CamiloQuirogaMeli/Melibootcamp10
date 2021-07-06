package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	older := 0
	for _, element := range employees {
		if element > 21 {
			older++
		}
	}
	fmt.Println("Empleados mayores a 21:", older)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println("Empleados finales: \n", employees)

}
