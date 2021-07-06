package main

import (
	"errors"
	"fmt"
	"os"
)

func validarSalario(salario int) (int, error) {
	if salario < 150000 {
		return salario, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return salario, nil
}
func main() {
	salary, err := validarSalario(10000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto ya que el salario es de %v", salary)
}
