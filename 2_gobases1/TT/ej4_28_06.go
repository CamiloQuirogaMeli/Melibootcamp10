package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ingrese el numero del mes")
	fmt.Println("---------------------")
	fmt.Print("-> ")
	var i int
	fmt.Scan(&i)

	switch i {
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
		fmt.Println("diciembre")
	default:
		fmt.Println("El numero no corresponde a ningun mes!")
	}

}
