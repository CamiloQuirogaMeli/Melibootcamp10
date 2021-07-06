package main

import (
	"fmt"
)

func main() {

	salary := 100000

	if salary < 150000 {
		err := fmt.Errorf("El mínimo es de 150000 y se ingresó: %d", salary)
		fmt.Print(err)
	} else {
		fmt.Print("Debe pagar impuestos")
	}
}
