package main

import (
	"fmt"
)

func main() {

	/*
		Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
		Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.

		var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

		Por otro lado también es necesario:
			saber cuántos de sus empleados son mayores a 21 años.
			Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
			Eliminar a Pedro del mapa.

	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var nuevoEmpleado string
	var edadNuevoEmpleado int
	var empleadoEliminado string
	var empleadosMayores []string

	//Imprimir
	ImprimirMap(employees, "Empleado", "Edad")

	//Contar mayores de 21 años
	empleadosMayores = ContarMayores(employees, 21)
	fmt.Println("Los empleados mayores a 21 años son:")
	imprimirSlice(empleadosMayores, "Empleado")
	fmt.Println("En total, hay", len(empleadosMayores), "empleados mayores a 21 años")

	//Agregar
	fmt.Println("Digita el nombre de un nuevo empleado(sugerencia: Federico):")
	fmt.Scanln(&nuevoEmpleado)

	fmt.Println("Digita la edad del nuevo empleado(sugerencia: 25):")
	fmt.Scanln(&edadNuevoEmpleado)

	employees[nuevoEmpleado] = edadNuevoEmpleado

	//Imprimir
	ImprimirMap(employees, "Empleado", "Edad")

	//Eliminar
	fmt.Println("Digita el nombre del empleado que se eliminará(sugerencia: Pedro):")
	fmt.Scanln(&empleadoEliminado)
	delete(employees, empleadoEliminado)

	//Imprimir
	ImprimirMap(employees, "Empleado", "Edad")
}

func ImprimirMap(mapa map[string]int, nombreClave string, nombreValor string) {
	i := 0
	for clave, valor := range mapa {
		i++
		fmt.Println(nombreClave, i, ":", clave, " -- ", nombreValor, ":", valor)
	}
}

func ContarMayores(mapa map[string]int, numero int) []string {

	var slice []string

	for clave, valor := range mapa {
		if valor > numero {
			slice = append(slice, clave)
		}
	}
	return slice
}

func imprimirSlice(slice []string, stringItem string) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(stringItem, i+1, ":", slice[i])
	}
}
