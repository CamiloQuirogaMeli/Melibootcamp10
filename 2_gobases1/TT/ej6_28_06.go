package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"], "años")

	fmt.Println("Mayores de 21 años:")
	for key, element := range employees {
		if element >= 21 {
			fmt.Println(key)
		}
	}

	fmt.Println("Agrego a Fede")
	employees["Federico"] = 25

	for key := range employees {
		fmt.Println(key)
	}

	fmt.Println("Elimino a Pedro")
	delete(employees, "Pedro")

	for key := range employees {
		fmt.Println(key)
	}
}
