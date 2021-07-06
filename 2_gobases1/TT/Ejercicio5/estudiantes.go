package main

import "fmt"

func main() {

	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
		"Leandro", "Eduardo", "Duvraschka"}

	var nombre string
	fmt.Print("Ingresar nombre del nuevo estudiante: ")
	fmt.Scan(&nombre)

	estudiantes = append(estudiantes, nombre)

	fmt.Println(estudiantes)

}
