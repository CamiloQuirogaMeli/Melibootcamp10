package main

import "fmt"

/*
Ejercicio 3 - Préstamo

Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

*/

func otorgaCredito(edad int, empleo bool, antiguedad int, sueldo float64) {

	/*Determino si se otorga o no el crédito en función de los parámetros*/

	if edad > 22 && empleo && antiguedad > 1 && sueldo > 100000.0 {

		fmt.Println("Se otorga crédito sin interés")

	} else if edad > 22 && empleo && antiguedad > 1 && !(sueldo > 100000.0) {

		fmt.Println("Se otorga crédito con interes")

	} else {

		fmt.Println("No se otorga crédito")

	}

}

func main() {

	edad := 23

	empleo := true

	antiguedad := 2

	sueldo := 100001.0

	fmt.Println("Datos del solicitante")
	fmt.Println("Edad: ", edad)
	fmt.Println("Empleo: ", empleo)
	fmt.Println("Antiguedad: ", antiguedad)
	fmt.Println("Sueldo: ", sueldo)

	otorgaCredito(edad, empleo, antiguedad, sueldo)

}
