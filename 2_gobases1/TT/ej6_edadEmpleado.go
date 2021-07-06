package main

import (
	"fmt"
)

func main() {
	var nameEmplo string
	var ageEmplo int
	var count int
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("De qué empleado desea saber la edad?")
	fmt.Scanln(&nameEmplo)
	fmt.Println("La edad de ", nameEmplo, "es: ", employees[nameEmplo])

	for _, k := range employees {
		if k > 21 {
			count++
		}
	}

	fmt.Println("Benjamin tiene: ", count, " empleados mayores de 21 años.")

	fmt.Println("Cuál es el nombre del nuevo empleado que desea ingresar a la lista?")
	fmt.Scanln(&nameEmplo)
	fmt.Println("Cuál es la edad de: ", nameEmplo)
	fmt.Scanln(&ageEmplo)

	employees[nameEmplo] = ageEmplo
	fmt.Println("La lista actualizada de empleados es: ", employees)

	fmt.Println("Cuál es el nombre del empleado que desea eliminar de la lista?")
	fmt.Scanln(&nameEmplo)
	delete(employees, nameEmplo)
	fmt.Println("La lista actualizada de empleados es: ", employees)
}
