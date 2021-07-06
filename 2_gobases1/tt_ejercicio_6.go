package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0

	fmt.Printf("Benjamin tiene %d años \n", employees["Benjamin"])

	for i, _ := range employees {
		if employees[i] > 21 {
			count++
		}
	}
	fmt.Printf("%d empelados son mayores de 21 años \n", count)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}
