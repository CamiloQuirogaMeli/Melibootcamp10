package main

import "fmt"

func main() {

	var salary float64 = 0
	var age int = 0
	var laborOld int = 0

	fmt.Printf("Ingrese su edad: ")
	fmt.Scanf("%d", &age)

	fmt.Printf("Ingrese su sueldo: ")
	fmt.Scanf("%f", &salary)

	fmt.Printf("Ingrese su antigüedad laboral en años: ")
	fmt.Scanf("%d", &laborOld)

	switch {
	case age >= 22:
		if salary > 100000 {
			if laborOld > 1 {
				fmt.Printf("Su prestamo será otorgado y se cobrara un interes")
			} else {
				fmt.Printf("No llega a la antigüedad necesaria")
			}
		} else {
			fmt.Printf("Su prestamo fue aprobado")
		}
	case age < 22:
		fmt.Printf("Su prestamo fue rechazado por no llegar a la edad mínima")

	}
}
