package main

import "fmt"

//Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
//Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que
//otorga no les cobrará interés a los que su sueldo es mejor a $100.000. Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a
//cada caso.

func main()  {
	var names [2] string
	var ages [2] int
	var emple [2] bool
	var tiemEmpl [2] float32
	var sal [2] int

	names[0] = "Leonardo"
	names[1] = "Juanita"

	ages[0] = 31
	ages[1] = 21

	emple[0] = true
	emple[1] = true

	tiemEmpl[0] = 10
	tiemEmpl[1] = 0.1

	sal[0] = 130000
	sal[1] = 70000

	for i := 0; i < len(names); i++ {
		if ages[i] > 22 {
			if emple[i] {
				if tiemEmpl[i] > 1 {
					fmt.Print("Felicitaciones!! El cliente " + names[i] + " recibe su préstamo...")
					if sal[i] > 100000 {
						fmt.Println("con cobro de intereses")
					} else {
						fmt.Println("sin cobro de intereses")
					}
				} else {
					fmt.Println("Lo sentimos!! El cliente " + names[i] + " no cumple con el tiempo mínimo como empleado")
				}
			} else {
				fmt.Println("Lo sentimos!! El cliente " + names[i] + " no es empleado")
			}
		} else {
			fmt.Println("Lo sentimos!! El cliente " + names[i] + " no cumple con la edad mínima")
		}
	}
}