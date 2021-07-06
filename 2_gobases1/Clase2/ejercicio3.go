package main

import (
	"fmt"
)

func main() {

	/*
		Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
		mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
		otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
		tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
		les cobrará interés a los que su sueldo es mejor a $100.000.
		Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
		acuerdo a cada caso.
	*/

	var edad float32
	var esEmpleadoYN string
	var esEmpleado bool
	var anhosEmpleoActual float32
	var sueldo float32
	var mensaje string

	fmt.Println("Digita las siguientes datos: ")
	//Edad
	fmt.Println("¿Cuantos años tienes?: ")
	fmt.Scanln(&edad)
	if edad > 22 {
		//Empleo
		fmt.Println("¿Tienes un empleo? (y/n)")
		fmt.Scanln(&esEmpleadoYN)
		//Validación empleo
		esEmpleado = validacionEmpleo(esEmpleadoYN)
		if esEmpleado == true {
			//Años de empleo
			fmt.Println("¿cuantos años llevas en tu empleo?")
			fmt.Scanln(&anhosEmpleoActual)
			if anhosEmpleoActual > 1 {
				mensaje = "Felicidades¡¡ se puede otorgar el crédito"
				//Sueldo
				fmt.Println("¿de cuanto es tu sueldo?")
				fmt.Scanln(&sueldo)
				if sueldo < 100000 {
					mensaje += " y por tu salario no se te cobrará interés"
				}
			} else {
				mensaje = "No se puede otorgar el crédito porque tienes que tener más de 1 año en tu empleo actual"
			}

		} else {
			mensaje = "No se puede otorgar el crédito porque tienes que tener un empleo"
		}
	} else {
		mensaje = "No se puede otorgar el crédito porque tienes que tener más de 22 años"
	}

	fmt.Println(mensaje)

}

func validacionEmpleo(esEmpleadoYN string) bool {
	for esEmpleadoYN != "y" && esEmpleadoYN != "Y" && esEmpleadoYN != "n" && esEmpleadoYN != "N" {
		fmt.Println("Dato no válido")

		fmt.Println("¿Tienes un empleo? (y/n)")
		fmt.Scanln(&esEmpleadoYN)
	}
	if esEmpleadoYN == "y" || esEmpleadoYN == "Y" {
		return true
	} else {
		return false
	}
}
