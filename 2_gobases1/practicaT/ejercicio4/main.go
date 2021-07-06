package main

import "fmt"

func main() {

	var months = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	var numberOfMonth int = 3

	fmt.Println(numberOfMonth, ",", months[numberOfMonth])

	/**
		Existen varias maneras de hacerlo, podriamos usar un array
		teniendo en cuenta que sabemos que es un numero constante de meses
		y dentro de cada posicion guardar el nombre del mes

		En este caso use un map por legibilidad de codigo
	**/

}
