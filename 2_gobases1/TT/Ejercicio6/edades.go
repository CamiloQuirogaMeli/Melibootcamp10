package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])
	var count int = 0
	for _, edad := range employees {
		if edad > 21 {
			count++
		}
	}
	fmt.Println("La cantidad de alumnos mayores a 21 años es: ", count)

	employees["Federico"] = 25

	fmt.Println("Lista con nuevo empleado: ", employees)

	delete(employees, "Pedro")

	fmt.Println("Lista sin Pedro: ", employees)

}
