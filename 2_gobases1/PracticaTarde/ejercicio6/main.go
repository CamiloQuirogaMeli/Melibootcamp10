package main

import "fmt"

func main() {
	var empleado string
	var edad int
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Ingrese nombre del empleado que desea consultar:")
	fmt.Scanln(&empleado)
	fmt.Println(empleado, "=>", employees[empleado])

	fmt.Println("Empleados mayores de 21:")
	for nombre, edad := range employees {
		if edad >= 21 {
			fmt.Println("Nombre: ", nombre, "=>", "Edad: ", edad)
		}
	}
	fmt.Println("Ingrese nombre del nuevo empleado: ")
	fmt.Scanln(&empleado)
	fmt.Println("Ingrese la edad del nuevo empleado: ")
	fmt.Scanln(&edad)

	employees[empleado] = edad

	fmt.Println("Nueva lista de empleados: ")
	for nombre, edad := range employees {

		fmt.Println("Nombre: ", nombre, "=>", "Edad: ", edad)

	}
	fmt.Println("Ingrese el nombre del empleado que desea eliminar: ")
	fmt.Scanln(&empleado)
	delete(employees, empleado)

	fmt.Println("Nueva lista de empleados: ")
	for nombre, edad := range employees {

		fmt.Println("Nombre: ", nombre, "=>", "Edad: ", edad)

	}
}
