package main

import (
	"fmt"
)

func main() {

	var edad uint
	var empleo string
	var antiguedad float64
	var sueldo float64

	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d \n", &edad)

	if edad > 21 {
		fmt.Println("Usted tiene empleo? s/n: ")
		fmt.Scanf("%v \n", &empleo)
		if empleo == "s" {
			fmt.Println("Ingrese los años de antigüedad: ")
			fmt.Scanf("%f.2 \n", &antiguedad)

			if antiguedad >= 1 {
				fmt.Println("Ingrese su sueldo: ")
				fmt.Scanf("%f.2 \n", &sueldo)

				if sueldo >= 100000 {
					fmt.Println("Felicitaciones! le otorgaremos un préstamo sin intereses!")
				} else {
					fmt.Println("Felicitaciones! le otorgaremos un préstamo con intereses!")
				}

			} else {
				fmt.Println("No puede acceder al prestamo. Es necesario que tenga mas de un año de antigüedad")
			}

		} else {
			fmt.Println("No puede acceder al prestamo. Es necesario que tenga empleo")
		}

	} else {
		fmt.Println("No puede acceder al prestamo. Es necesario ser mayor de 22 años")
	}

}
