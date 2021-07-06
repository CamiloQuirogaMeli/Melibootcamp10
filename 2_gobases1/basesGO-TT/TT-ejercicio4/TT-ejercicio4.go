package main

import "fmt"

/*

Ejercicio 4 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
por qué?
Ej: 7, Julio

*/

func imprimeMes(mes int) {

	/*Encuentro dos maneras de resolver el ejercicio, empleando la estructura switch y un conjunto de if.
	Opto por la estructura switch debido a que me proporciona un codico más legible y organizado*/

	switch mes {
	case 1:
		fmt.Println(mes, ", Enero")
	case 2:
		fmt.Println(mes, ", Febrero")
	case 3:
		fmt.Println(mes, ", Marzo")
	case 4:
		fmt.Println(mes, ", Abril")
	case 5:
		fmt.Println(mes, ", Mayo")
	case 6:
		fmt.Println(mes, ", Junio")
	case 7:
		fmt.Println(mes, ", Julio")
	case 8:
		fmt.Println(mes, ", Agosto")
	case 9:
		fmt.Println(mes, ", Septiembre")
	case 10:
		fmt.Println(mes, ", Octubre")
	case 11:
		fmt.Println(mes, ", Noviembre")
	case 12:
		fmt.Println(mes, ", Diciembre")
	default:
		fmt.Println("Valor desconocido")
	}

}

func main() {

	mes := 9

	imprimeMes(mes)

}
