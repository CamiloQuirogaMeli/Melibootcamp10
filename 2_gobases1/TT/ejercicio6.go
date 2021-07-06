package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// edad de Benjamin
	fmt.Println("La edad de Benjamin es de:", employees["Benjamin"], "años")

	// empleados mayores a 21 años
	var employeesTwentyOneYearsOld int

	for _, employee := range employees {
		if employee > 21 {
			employeesTwentyOneYearsOld++
		}
	}
	fmt.Println(employeesTwentyOneYearsOld, "empleados tienen mas de 21 años")

	// agregar a Federico de 25 años
	employees["Federico"] = 25
	fmt.Println(employees)

	// eliminar a Pedro de la lista de empleados
	delete(employees, "Pedro")
	fmt.Println(employees)
}
