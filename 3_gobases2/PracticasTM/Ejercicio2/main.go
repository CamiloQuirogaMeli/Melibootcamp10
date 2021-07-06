package main

import (
	"fmt"
)

func main() {

	imp := promedio(5, 5, 5, 5, 5)
	if imp == 0 {
		fmt.Println("No se pudo calcular promedio")
	} else {
		fmt.Println("El promedio del estudiante es: ", imp)
	}

}

func promedio(promedio ...float64) float64 {
	var suma float64 = 0

	for i := 0; i < len(promedio); i++ {
		if promedio[i] < 0 {
			fmt.Println("Hay un error con las notas, verifique las notas, no se puede calcular el promedio")
			return 0
		}
		suma += promedio[i]
	}
	div := float64(len(promedio))

	return suma / div
}
