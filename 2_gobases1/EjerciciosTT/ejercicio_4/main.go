package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Ingrese el numero de mes: ")
	var mes int
	fmt.Scanln(&mes)

	a := time.Month(mes)

	fmt.Println("El mes ingresado es: ", a)

}
