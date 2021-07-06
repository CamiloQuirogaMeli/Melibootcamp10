// Ejercicio 3 - Préstamo

// Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
// mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
// otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
// tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
// les cobrará interés a los que su sueldo es mejor a $100.000.
// Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
// acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var age int
	var employed string
	var antique int
	var salary float64

	fmt.Printf("Ingresar la edad: ")
	fmt.Scan(&age)
	fmt.Printf("¿Se encuentra empleado?: ")
	fmt.Scan(&employed)
	fmt.Printf("Ingrese la antiguedad: ")
	fmt.Scan(&antique)
	fmt.Printf("Ingrese su salario: ")
	fmt.Scan(&salary)

	employed = strings.ToLower(employed)

	if strings.Compare(employed, "si") == 0 && antique > 1 && age > 22 {
		if salary > 100000 {
			fmt.Printf("Se le cobraran intereses")
		} else {
			fmt.Printf("No se le cobraran intereses")
		}
	} else if age <= 22 {
		fmt.Printf("No posee la edad necesaria")
	} else if strings.Contains(employed, "no") {
		fmt.Print("No se encuentra empleado, por lo que no puede recibir un prestamo")
	} else if antique <= 1 {
		fmt.Printf("Su antiguedad no es suficiente")
	}

}
