package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//Item A
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	//Item B
	empMayores := 0
	for _, element := range employees {
		if element > 21 {
			empMayores++
		}
	}
	fmt.Println("Mayores a 21: ", empMayores)

	//Item C
	employees["Federico"] = 25

	//Item D
	delete(employees, "Pedro")

}
