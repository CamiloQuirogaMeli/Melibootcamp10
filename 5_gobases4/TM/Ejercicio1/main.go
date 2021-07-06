package main

import (
	"fmt"
	"os"
)

type errorPersonalizado struct {
	mensaje string
	codigo  int
}

func (e *errorPersonalizado) Error() string {
	return fmt.Sprintf("%d - %v", e.codigo, e.mensaje)
}

func errorPersonalizadoTest(sueldo int) error {
	if sueldo < 150000 {
		return &errorPersonalizado{
			mensaje: "error: el salario ingresado no alcanza el mínimo no imponible",
			codigo:  400,
		}
	}
	return nil
}

func main() {
	var salary int = 125000

	err := errorPersonalizadoTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
