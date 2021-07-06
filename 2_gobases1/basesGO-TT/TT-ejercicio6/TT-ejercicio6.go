package main

import "fmt"

/*

Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19,
"Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
- Saber cuántos de sus empleados son mayores a 21 años.
- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
- Eliminar a Pedro del mapa.

*/

/*Listo empleados*/
func listaEmpleados(empleados map[string]int) {

	for empleado, edad := range empleados {

		fmt.Println(empleado, edad)

	}

}

/*Muestro la edad del empleado*/
func edadEmpleado(empleados map[string]int, clave string) {

	fmt.Println("Edad de ", clave, ": ", empleados[clave])

}

/*Empleados mayores de 21*/
func empleadosMayores(empleados map[string]int) {

	mayores21 := 0
	for _, edad := range empleados {

		if edad > 21 {
			mayores21++
		}

	}
	fmt.Println("Cantidad de empleados mayores a 21: ", mayores21)

}

/*Agrego una nueva entrada al mapa*/
func agregaEmpleado(empleados map[string]int, clave string, edad int) {

	empleados[clave] = edad
	fmt.Println("Empleados actualizados: ")
	listaEmpleados(empleados)

}

/* Elimino una clave y su valor */
func eliminaEmpleado(empleados map[string]int, clave string) {

	delete(empleados, clave)
	fmt.Println("Empleados actualizados: ")
	listaEmpleados(empleados)

}

func main() {

	var empleados = map[string]int{"Benjamín": 20, "Nahuel": 26, "Brenda": 19,
		"Darío": 44, "Pedro": 30}

	/*Listo empleados*/
	listaEmpleados(empleados)

	/*Obtengo la edad de Benjamín*/
	var clave string = "Benjamín"
	edadEmpleado(empleados, clave)

	/*Empleados mayores de 21*/
	empleadosMayores(empleados)

	/*Agrego nuevo empleado y muestro lista actualizada*/
	nombre := "Federico"
	edad := 25
	agregaEmpleado(empleados, nombre, edad)

	/*Elimino empleado y muestro lista actualizada*/
	name := "Pedro"
	eliminaEmpleado(empleados, name)

}
