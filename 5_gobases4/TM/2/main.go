package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int

	fmt.Println("Ingrese el sueldo:")
	fmt.Scanln(&salary)

	if salary < 150000 {
		err := errors.New("error: el salario ingresado no alcanza el minimo imponible")
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
