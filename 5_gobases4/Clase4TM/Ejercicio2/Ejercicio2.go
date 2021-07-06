package main

import (
	"errors"
	"fmt"
)

func main() {

	salary := 160000

	if salary < 150.000 {
		err := errors.New("El salario ingresado no alcanza el minimo")
		fmt.Print(err)
	} else {
		fmt.Print("Debe pagar impuestos")
	}
}
