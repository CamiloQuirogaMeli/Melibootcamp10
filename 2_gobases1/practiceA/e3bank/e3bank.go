// Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
// mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
// otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
// tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
// les cobrará interés a los que su sueldo es mejor a $100.000.
// Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
// acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main

import "fmt"

func main() {
	var (
		age    uint64
		job    float64
		salary float64
	)

	fmt.Printf("Ingrese edad: ")
	fmt.Scanf("%d\n", &age)

	if age >= 22 {
		fmt.Printf("Ingrese la antiguedad laboral en meses: ")
		fmt.Scanf("%f\n", &job)
		if job > 12 {
			fmt.Printf("Ingrese salario: ")
			fmt.Scanf("%f\n", &salary)
			if salary > 100000 {
				fmt.Println("Apto para préstamos sin intereses.")
			} else {
				fmt.Println("Apto para prestamo.")
			}
		} else {
			fmt.Println("No se otorgan préstamos a quienes no tengan al menos un año de antiguedad en su trabajo")
		}
	} else {
		fmt.Println("No se otorgan préstamos a menores de 22 años.")
	}

}
