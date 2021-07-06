package main

import (
	"fmt"
	"strings"
)

/*Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
acuerdo a cada caso.
le da prestamos a: mayores de 22, empleados hace mas 1 año.
Si cobra mas de 100 mil, no cobra intereses*/
func main() {
	var edad uint8
	var empleado string
	var experiencia uint8
	var sueldo uint
	const EDADMINIMA uint8 = 22
	const SUELDOSININT uint = 100000

	fmt.Println("Chequeo de aplicacion a prestamos:")
	fmt.Println("Ingresar edad del solicitante:")
	fmt.Scanln(&edad)
	if edad < EDADMINIMA {
		fmt.Println("Debe ser mayor de 22 anios para solicitar un prestamo")
	} else {
		fmt.Println("Actualmente esta empleado? si/no")
		fmt.Scanln(&empleado)
		empleado = strings.ToLower(empleado)
		if empleado != "si" {
			fmt.Println("Debe estar empleado para solicitar un prestamo")
		} else {
			fmt.Println("Cuantos años tiene de experiencia laboral? Indique 0 si es menor a 1 año")
			fmt.Scanln(&experiencia)
			if experiencia < 1 {
				fmt.Println("Debe tener mas de 1 anio de experiencia para solicitar un prestamo")
			} else {
				fmt.Println("Ingrese el sueldo del solicitante: ")
				fmt.Scanln(&sueldo)
				if sueldo >= SUELDOSININT {
					fmt.Println("Se le puede aplicar prestamo sin intereses")
				} else {
					fmt.Println("Se le puede aplicar prestamo con intereses")
				}
			}
		}
	}

}
