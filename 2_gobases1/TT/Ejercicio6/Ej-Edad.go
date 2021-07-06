package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cont uint16 = 0

	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	for key, element := range employees {
		if key != "" && element > 21 {
			cont++
		}
	}
	fmt.Println("La cantidad de empleados mayores de 21 años es: ", cont)

	employees["Federico"] = 25

	fmt.Println("Federico fue agregado a la lista. Lista:\n", employees)

	delete(employees, "Pedro")

	fmt.Println("Pedro fue eliminado de la lista. Lista:\n", employees)
}

// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
// Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.

// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// - Saber cuántos de sus empleados son mayores a 21 años.

// 3

// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// - Eliminar a Pedro del mapa.
