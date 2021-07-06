package main

import "fmt"

/*
	Ejercicio 6 - Que edad tiene ...
		Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
		Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.
		Por otro lado también es necesario:
			- Saber cuántos de sus empleados son mayores a 21 años.
			- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
			- Eliminar a Pedro del mapa.
*/

func main() {
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println(empleados)                                        // Muestra los empleados
	fmt.Println("La edad de benjamin es:", empleados["Benjamin"]) // Trae la edad del empleado Benjamin
	fmt.Println("Los empleados mayores a 21 años son:")
	for nombre, edad := range empleados {
		if edad > 21 {
			fmt.Println("Nombre", nombre, "=>", "Edad", (edad)) //  Muestra los empleados mayores a 21 años
		}
	}
	empleados["Federico"] = 25 // Añade un nuevo emplado
	fmt.Println(empleados)     // Muestra los empleados
	delete(empleados, "Pedro") // Elimina de la lista al empleado pedro
	fmt.Println(empleados)     // Muestra los empleados
}
