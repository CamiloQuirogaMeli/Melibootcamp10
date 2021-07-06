package main

import (
	"fmt"
	"strings"
)

func main() {
	var sueldo float32
	var edad, anios int
	var trabaja string
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)

	if edad >= 22 {
		fmt.Println("Trabaja? (s/n): ")
		fmt.Scanln(&trabaja)

		if strings.ToLower(trabaja) == "s" {
			fmt.Println("¿Cuantos años lleva trabajando?: ")
			fmt.Scanln(&anios)

			if anios > 1 {
				fmt.Println("Usted está habilitado a solicitar un prestamo")

				fmt.Println("Ingrese su sueldo: ")
				fmt.Scanln(&sueldo)

				if sueldo > 100000 {
					fmt.Println("No se le cobrarán intereses")
				} else {
					fmt.Println("Se le cobrarán intereses")
				}
			} else {
				fmt.Println("Debe tener más de un año de antiguedad para obtener el prestamo")
			}
		} else {
			fmt.Println("Debe estar empleado para obtener un prestamo")
		}
	} else {
		fmt.Println("Debe ser mayor de 21 años para obtener un prestamo")
	}
}
