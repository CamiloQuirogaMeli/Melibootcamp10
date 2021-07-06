package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 3")

	var age, antiquityMonth int = 23, 14
	var isEmployee bool = true
	var salary float32 = 100000
	if age > 22 {
		if isEmployee {
			if antiquityMonth > 12 {
				fmt.Println("Usted es elegible para un prestamo")
				if salary > 100000.00 {
					fmt.Println("Ademas su prestamo no tendra interes")
				} else {
					fmt.Println("Su prestamo tendra interes debido a su salario")
				}
			} else {
				fmt.Println("Necesita al menos 1 año para acceder a un prestamo")
			}
		} else {
			fmt.Println("La persona debe ser empleado/a para que se le otorgue un prestamo")
		}
	} else {
		fmt.Println("El prestamo no puede ser otorgado a menores de 22")
	}

}
