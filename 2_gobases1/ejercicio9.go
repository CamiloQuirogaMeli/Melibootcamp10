package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func ejercicio9() {
	var count int = 0
	fmt.Println(employees["Benjamin"])

	for employee := range employees {
		if employees[employee] > 21 {
			count++
		}
	}
	delete(employees, "Pedro")
	employees["Federico"] = 25
	fmt.Printf("Cantidad de empleados mayores a 21 años: %v", count)
}
