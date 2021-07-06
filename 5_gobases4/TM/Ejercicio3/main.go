package main

import (
	"fmt"
	"os"
)

func validarSalario(salario int) (int, error) {
	if salario < 150000 {
		return salario, fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: %v", salario)
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
