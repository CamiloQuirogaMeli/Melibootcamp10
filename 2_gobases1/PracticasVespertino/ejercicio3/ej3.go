package main

import (
	"fmt"
)

type Person struct {
	name       string
	age        int
	experience int
	salary     float64
	active     bool
}

func main() {

	person := Person{"Juan Martin", 23, 3, 100.5, true}

	available := isAvailable(person)

	if available {
		fmt.Printf("Ej usuario %s se encuentra habilitado para recibir un prestamo\n", person.name)

		getTaxes := getTaxes(person.salary)

		if getTaxes {
			fmt.Println("El prestamo recibira impuestos")
		} else {
			fmt.Println("El prestamo no recibira impuestos")
		}

	} else {
		fmt.Printf("%s no se encuentra habiltado para recibir un prestamo\n", person.name)
	}

}

func isAvailable(person Person) bool {

	if person.age > 22 && person.active && person.experience > 1 {
		return true
	}

	return false
}

func getTaxes(salary float64) bool {

	return salary <= 1000.0
}

/*
	 Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
	mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.

	 Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
	tengan más de un año de antigüedad en su trabajo.

	 Dentro de los préstamos que otorga no
	les cobrará interés a los que su sueldo es mejor a $100.000.

	Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
	acuerdo a cada caso.

	Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/
