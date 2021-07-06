package main

import (
	"errors"
	"fmt"
)

func main() {

	salary := 160000

	if salary < 150.000 {
		e := errors.New("error: el salario ingresado no alcanza el minimo imponible")
		fmt.Print(e)
	} else {
		fmt.Print("Debe pagar impuestos")
	}
}
