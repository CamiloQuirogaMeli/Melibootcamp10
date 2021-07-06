package main

import (
	"fmt"
)

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var (
		empleadosMayores int
	)

	empleadosMayores = 0

	// edad
	fmt.Println("Edad de Benjamin:", employees["Benjamin"])

	for _, element := range employees {
		if element > 21 {
			empleadosMayores++
		}
	}

	fmt.Printf("Hay %d empleados mayores a 21\n", empleadosMayores)

	// agrego a federico
	employees["Federico"] = 25
	fmt.Print(employees)

	// elimino a Pedro
	delete(employees, "Pedro")
	fmt.Print(employees)
}
