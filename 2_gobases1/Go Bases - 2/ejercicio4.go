package main

import "fmt"

/*
	Ejercicio 4 - A que mes corresponde
		Realizar una aplicación que contenga una variable con el número del mes, imprimir
		el mes que corresponda. ¿Se te ocurre si se puede resolver de más de una manera?
		¿Cuál elegirías y por qué? Ej: 7, Julio

		Rta : Tambien  se puede hacer con un if else pero seria un proceso mas largo en
		codigo el switch es mas corto y especifico
*/

func main() {
	numeroMes := 6
	fmt.Println("El numero es el", numeroMes)
	switch numeroMes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Numero no corresponde a ningun mes")
	}
}
