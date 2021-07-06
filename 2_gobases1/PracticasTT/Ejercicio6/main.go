package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])
	var suma int = 0
	for _, employee := range employees {
		if employee > 21 {
			suma += 1

		}

	}
	fmt.Println("Hay ", suma, " empleado mayores de 21 años")
	employees["Federico"] = 25
	fmt.Println("Se agrego un nuevo empleado de 25 años, la lista nueva seria ", employees)
	delete(employees, "Pedro")
	fmt.Println("Se elimino el empleado Pedro, la nueva lista quedaria ", employees)
}
