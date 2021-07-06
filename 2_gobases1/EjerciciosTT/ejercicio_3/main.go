package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ingrese su edad: ")
	var edad int
	fmt.Scanln(&edad)

	fmt.Println("Tiene trabajo actualemnte? (s/n): ")
	var conEmpleo string
	fmt.Scanln(&conEmpleo)

	fmt.Println("Indique antiguedad (en meses): ")
	var antiguedad int
	fmt.Scanln(&antiguedad)

	fmt.Println("Ingrese su sueldo: ")
	var sueldo int
	fmt.Scanln(&sueldo)

	fmt.Println("Edad: ", edad)
	fmt.Println("Con empleo: ", conEmpleo)
	fmt.Println("Antiguedad: ", antiguedad)

	if edad > 22 && conEmpleo == "s" && antiguedad > 12 {
		fmt.Println("El banco le otorgará un prestamo")

		if sueldo > 100000 {
			fmt.Println("Y NO le cobrará intereses")
		} else {
			fmt.Println("Y le cobrará intereses")
		}
	} else {
		fmt.Println("El banco NO le otorgará un prestamo")
	}

}
