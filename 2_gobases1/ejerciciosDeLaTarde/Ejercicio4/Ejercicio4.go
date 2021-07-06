package main

import "fmt"

func main() {

	var month int
	fmt.Printf("Ingrese el mes a comprobar:")
	fmt.Scanf("%d", &month)

	switch month {
	case 1:
		fmt.Printf("Enero")
	case 2:
		fmt.Printf("Febrero")
	case 3:
		fmt.Printf("Marzo")
	case 4:
		fmt.Printf("Abril")
	case 5:
		fmt.Printf("Mayo")
	case 6:
		fmt.Printf("Junio")
	case 7:
		fmt.Printf("Julio")
	case 8:
		fmt.Printf("Agosto")
	case 9:
		fmt.Printf("Septiembre")
	case 10:
		fmt.Printf("Octubre")
	case 11:
		fmt.Printf("Noviembre")
	case 12:
		fmt.Printf("Diciembre")
	default:
		fmt.Printf("Numero invalido")
	}

}
