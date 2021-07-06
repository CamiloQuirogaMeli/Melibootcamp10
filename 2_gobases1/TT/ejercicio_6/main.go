package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	// a) --------------------------
	fmt.Printf("Benjamin tiene %d años\n", employees["Benjamin"])

	// b) --------------------------
	fmt.Printf("La cantidad de empleados mayores a 21 es %d\n", employeesOverAge(employees, 21))

	// c) --------------------------
	employees["Federico"] = 25
	fmt.Println("Empleados luego de agregar a Federico:")
	printEmployees(employees)

	// d) --------------------------
	delete(employees, "Pedro")
	fmt.Println("Empleados luego de sacar a Pedro:")
	printEmployees(employees)
}

func employeesOverAge(employees map[string]int, age int) int {
	var count int = 0
	for _, value := range employees {
		if value > age {
			count++
		}
	}
	return count
}

func printEmployees(employees map[string]int) {
	for name, age := range employees {
		fmt.Printf("Nombre: %s, Edad: %d\n", name, age)
	}
}
