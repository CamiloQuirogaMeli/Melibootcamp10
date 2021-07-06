package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad de Benjamin", employees["Benjamin"])

	counter := 0
	for _, element := range employees {
		if element > 21 {
			counter++
		}
	}
	fmt.Println("Cantidad de empleados mayores a 21:", counter)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
