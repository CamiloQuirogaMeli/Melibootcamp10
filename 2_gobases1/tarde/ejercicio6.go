package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var BenjaminsAge = employees["Benjamin"]
	fmt.Println("La edad del empleado Benjamin es", BenjaminsAge)
	fmt.Println("Número de empleados con más de 20 años:", employeesOver21(employees))

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Lista de empleados luego de ingresar a Federico y sacar a Pedro")
	fmt.Println(employees)
}

func employeesOver21(employees map[string]int) int {
	var employeesOver21 int = 0
	for _, age := range employees {
		if age > 20 {
			employeesOver21++
		}
	}
	return employeesOver21
}
