package main

import f "fmt"
	
func main() {
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

	var age uint 
	var employee string
	var antiquity uint
	var salary uint

	f.Println("Ingrese su edad")
	f.Scanln(&age)
	f.Println("Es empleado? Si/No")
	f.Scanln(&employee)
	f.Println("Ingrese su antiguedad")
	f.Scanln(&antiquity)
	f.Println("Ingrese su sueldo")
	f.Scanln(&salary)

	if age >= 22 {
		if employee == "Si" &&  antiquity >= 1{
			f.Println("Yes! Tendras un prestamo!")
			if salary >= 100000 {
				f.Println("Y no se te cobraran interes ;)") 
			}else{
				f.Println("Te informaremos de los intereses") 
			}
		}
	}else {
		f.Println("Lo siento, eres muy joven para esto :c")
	}
}
