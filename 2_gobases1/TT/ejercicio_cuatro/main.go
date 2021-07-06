package main

import (
	"fmt"
)

func main() {

	var (
		nroMes    int
		stringMes string
	)

	fmt.Print("ingrese numero del mes: ")
	fmt.Scanf("%d", &nroMes)

	if nroMes < 1 || nroMes > 12 {
		fmt.Print("ingreso un numero incorrecto")
	} else {
		switch nroMes {
		case 1:
			stringMes = "Enero"
		case 2:
			stringMes = "Febrero"
		case 3:
			stringMes = "Marzo"
		case 4:
			stringMes = "Abril"
		case 5:
			stringMes = "Mayo"
		case 6:
			stringMes = "Junio"
		case 7:
			stringMes = "Julio"
		case 8:
			stringMes = "Agosto"
		case 9:
			stringMes = "Septiembre"
		case 10:
			stringMes = "Octubre"
		case 11:
			stringMes = "Noviembre"
		case 12:
			stringMes = "Diciembre"
		default:
			stringMes = "Error"
		}
		fmt.Println("Estamos en", stringMes)
	}

}

// Comentario: Se me ocurrio resolverlo mediante el ingreso estandar del numero de mes, validar a que nombre de mes corresponde e imprimirlo como tal.
