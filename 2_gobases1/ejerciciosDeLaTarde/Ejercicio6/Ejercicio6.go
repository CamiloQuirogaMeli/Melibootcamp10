package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var quantity int
	fmt.Println(employees["Benjamin"])

	for nombre, edad := range employees {
		if edad < 21 {
			fmt.Println("Nombre: ", nombre, " => Edad: ", edad)
			quantity++
		}
	}
	fmt.Println("los empleados menores de 21 son: ", quantity)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println("Se elimina a Pedro del mapa ")
	for nombre, edad := range employees {
		fmt.Println("Nombre: ", nombre, " => Edad: ", edad)
	}
}

//Imprimir la edad de Benjamin
//Por otro lado también es necesario:
//- Saber cuántos de sus empleados son mayores a 21 años.
//3
//- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
//- Eliminar a Pedro del mapa.
