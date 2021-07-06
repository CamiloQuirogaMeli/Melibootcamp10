package main

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

type myError struct {
	montoSalario int
	msg          string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.montoSalario, e.msg)
}

func validarSalario(salario int) (int, error) {
	if salario < 150000 {
		return salario, &myError{
			montoSalario: salario,
			msg:          "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return salario, nil
}
func main() {
	salary, err := validarSalario(160000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto ya que el salario es de %v", salary)
}
