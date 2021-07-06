package main

import (
	"fmt"
)

func main() {

	var mes uint

	fmt.Println("Ingrese un número entre el 1 y el 12: ")
	fmt.Scanf("%v \n", &mes)

	switch mes {
	case 1:
		fmt.Println("El número corresponde al mes de Enero")
	}

}
