package main

import f "fmt"

func main() {
	/*
	Ejercicio 6 - Qué edad tiene...
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	Según el siguiente mapa, ayuda a imprimir la edad de Benjamin.
	var employees = map[string]int{"Benjamin":20, "Nahuel":26, "Brenda":19, "Dario":44, "Pedro":30}
	Por otro lado también es necesario:
	- Saber cuántos de sus empleados son mayores a 21 años.
	- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	- Eliminar a Pedro del mapa.
	*/

	var employees = map[string]int{"Benjamin":20, "Nahuel":26, "Brenda":19, "Dario":44, "Pedro":30}

	f.Println("Edad de Benjamin:", employees["Benjamin"])

	cantAge := 0
	
	for _, element := range employees {
		if element > 21 {
			cantAge++
		}
	}
	f.Println("Empleados mayores a 21 años: ", cantAge)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	f.Println("Empleados finales")
	for key, element := range employees {
		f.Println(key, element, "años")
	}

}
