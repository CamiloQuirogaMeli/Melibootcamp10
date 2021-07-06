package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 125000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo no imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}
