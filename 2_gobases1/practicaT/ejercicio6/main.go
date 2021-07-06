package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	var name string = "Benjamin"
	fmt.Println("Nombre: ", name)
	fmt.Println("Edad: ", employees[name])

	var upTwentyOne = 0
	for _, employee := range employees {
		if int(employee) > 21 {
			upTwentyOne++
		}
	}

	fmt.Println("Hay ", upTwentyOne, " empleados mayores de 21")
	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)

}
