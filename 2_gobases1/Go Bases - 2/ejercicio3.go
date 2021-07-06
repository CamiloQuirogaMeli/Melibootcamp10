package main

import "fmt"

/*
	Ejercicio 3 - Prestamo
		Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
		Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga
		préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de
		un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés
		a los que su sueldo es mejor a $100.000.
		Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
		acuerdo a cada caso.

		Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

func main() {
	var edad, interes int = 0, 3
	var empleado, experiencia, sueldo bool = false, false, false

	fmt.Print("Introduzca su edad: ")
	fmt.Scan(&edad)
	if edad > 22 {
		fmt.Print("Actualmente esta trabajando 0:NO/1:SI ")
		fmt.Scan(&empleado)
		if empleado == true {
			fmt.Print("Tienes mas de un año de experiencia en tu actual trabajo 0:NO/1:SI ")
			fmt.Scan(&experiencia)
			if experiencia == true {
				fmt.Print("Tu sueldo es mayor a $100.000 0:NO/1:SI ")
				fmt.Scan(&sueldo)
				if sueldo == true {
					fmt.Println("Se aprobo tu prestamo, adicional por ser tu sueldo mayor a $100.000 no cobraremos interes en tu prestamo")
				} else {
					fmt.Printf("Se aprobo tu prestamo, con un interes del %d porciento\n", interes)
				}
			} else {
				fmt.Println("No cuentas con la experiencia necesaria para solicitar un prestamo")
			}
		} else {
			fmt.Println("No estas actualmente activo como trabajador para solicitar un prestamo")
		}
	} else {
		fmt.Println("No cumples la mayoria de edad para solicitar un prestamo")
	}
}
