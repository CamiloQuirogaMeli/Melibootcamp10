package main

import "fmt"

func main() {

	//
	// Ejercicio 3 - Préstamo
	// Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
	// mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
	// otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
	// tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
	// les cobrará interés a los que su sueldo es mejor a $100.000.
	// Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
	// acuerdo a cada caso.
	//
	// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
	//

	var (
		edad            = 25
		antiguedadAnios = 2
		sueldo          = 101000
		esEmpleado      = true
	)

	rspEsEmpleado := (map[bool]string{true: "Sí", false: "No"})

	fmt.Println("===== Datos del cliente: =====")
	fmt.Println("1) Es empleado:", rspEsEmpleado[esEmpleado])
	fmt.Println("2) Edad:", edad)
	fmt.Printf("3) Antigüedad: %d año(s)\n", antiguedadAnios)
	fmt.Printf("4) Sueldo: $%d\n", sueldo)

	fmt.Println("===== Respuesta del Banco: =====")

	switch {
	case esEmpleado && edad > 22 && antiguedadAnios > 1 && sueldo > 100000:
		fmt.Println("Felicidades cumples con los requisitos para adquirir un préstamo,\nademás por tener un sueldo superior a $100.000 no te cobraremos intereses")
	case esEmpleado && edad > 22 && antiguedadAnios > 1:
		fmt.Println("Felicidades cumples con los requisitos para adquirir un préstamo")
	default:
		fmt.Println("Lo sentimos no puedes adquirir un préstamo")
	}

}
