package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin tiene", employees["Benjamin"], "años")

	var olderThan int = 0

	for _, value := range employees {
		if value > 25 {
			olderThan += 1
		}
	}

	fmt.Println("Hay", olderThan, "empleados mayores de 25")

	employees["Federico"] = 25
	fmt.Println(employees, "-> Federico añadido")

	delete(employees, "Pedro")
	fmt.Println(employees, "-> Adiós, Pedro :C")
}
