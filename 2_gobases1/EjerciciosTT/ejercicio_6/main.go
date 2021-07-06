package main

import (
	"fmt"
)

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Empleados: ", employees)

	fmt.Println(employees["Benjamin"])

	var mayores int
	for _, employee := range employees {
		// fmt.Println("key: ", key, " -> ", "Value: ", employee)
		if employee > 21 {
			mayores++
		}
	}
	fmt.Println("Hay ", mayores, " empleados mayores a 21 años")

	employees["Federico"] = 25

	fmt.Println("Se agrega a Federico: ", employees)

	delete(employees, "Pedro")

	fmt.Println("Se elimina a Pedro: ", employees)
}
