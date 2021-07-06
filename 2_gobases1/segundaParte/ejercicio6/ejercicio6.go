package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es %d\n", employees["Benjamin"])
	var ageMajority int
	for _, val := range employees {
		if val >= 21 {
			ageMajority++
		}
	}
	fmt.Printf("La cantidad de empleados mayores a 21 son %d\n", ageMajority)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
