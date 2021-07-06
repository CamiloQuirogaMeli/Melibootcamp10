// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
// Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.

// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// - Saber cuántos de sus empleados son mayores a 21 años.

// 3

// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// - Eliminar a Pedro del mapa.

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var contEmployees int = 0

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	for key, element := range employees {
		fmt.Printf("Empleado: %s , Edad: %d \n", key, element)
		if element > 21 {
			contEmployees += 1
		}
	}

	fmt.Println("Se agrega un nuevo empleado Federico de 25 años")

	employees["Federico"] = 25

	fmt.Println("Se elimina al empleado Pedro")

	delete(employees, "Pedro")

	fmt.Println("La cantidad de empleados mayores a 21 es:", contEmployees)

	fmt.Println("Los nuevos empleados son: ")
	for key, element := range employees {
		fmt.Printf("Empleado: %s , Edad: %d \n", key, element)
	}

}
