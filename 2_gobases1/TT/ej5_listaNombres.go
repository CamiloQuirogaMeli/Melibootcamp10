package main

import (
	"fmt"
)

func listaNom() {
	var nameSlice []string
	var ask string
	var newName string
	nameSlice = append(nameSlice, "Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka")

	fmt.Println("Desea ingresar un nuevo alumno? Y (si) o N (no)")
	fmt.Scanln(&ask)

	if ask == "Y" {
		fmt.Println("Cómo es el nombre del nuevo alumno?: ")
		fmt.Scanln(&newName)
		nameSlice = append(nameSlice, newName)
		fmt.Println("La lista actualizada es: ", nameSlice)
	} else {
		fmt.Println("La lista es: ", nameSlice)
	}
}
