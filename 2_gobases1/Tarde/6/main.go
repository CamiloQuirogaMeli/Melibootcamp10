package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	numberOlder := 0

	fmt.Println("Edad de Benjanmin :", employees["Benjamin"])

	for _, element := range employees {
		if element > 21 {
			numberOlder++
		}
	}

	fmt.Println("Cantidad de empleados mayores a 21 años :", numberOlder)
	employees["Federico"] = 22
	delete(employees, "Pedro")

}
