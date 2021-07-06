package main

import (
	"fmt"
)

func main() {

	lista := []string{"Benjamin,", "Nahuel,", "Brenda,", "Marcos,", "Pedro,", "Axel,", "Alez,", "Dolores,", "Federico,", "Hernan,", "Leandro,", "Eduardo,", "Duvraschka"}

	fmt.Println("El Profesor tiene la siguiente lista de estudiantes: ", lista, "\n\n luego de dos clases se sumo una chica\n\n ")
	lista = append(lista, ", Gabriela")
	fmt.Println("por lo cual la nueva lista se compone de: ", lista)

}
