package main

import (
	"fmt"
)

func main() {

	salary := 160000

	if salary < 150.000 {
		e := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Print(e)
	} else {
		fmt.Print("Debe pagar impuestos")
	}
}
