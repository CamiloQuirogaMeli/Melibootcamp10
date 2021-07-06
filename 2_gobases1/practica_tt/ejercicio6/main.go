package main

import "fmt"

func main() {

	//
	// Ejercicio 6 - Qué edad tiene...
	// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	// Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.
	// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	//
	// Por otro lado también es necesario:
	// - Saber cuántos de sus empleados son mayores a 21 años.
	// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	// - Eliminar a Pedro del mapa.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	// Número de empleados con edad mayor a 21 años
	numEmployees := 0
	for _, edad := range employees {
		if edad > 21 {
			numEmployees++
		}
	}
	fmt.Println("Total de empleados mayores a 21 años:", numEmployees)

	// Agregamos un nuevo empleado
	fmt.Println("Listado de empleados")
	fmt.Println(employees)
	employees["Federico"] = 25
	fmt.Println("Listado de empleados después de agregar a Federico")
	fmt.Println(employees)

	// Eliminamos a Pedro
	delete(employees, "Pedro")
	fmt.Println("Listado de empleados después de eliminar a Pedro")
	fmt.Println(employees)

}
