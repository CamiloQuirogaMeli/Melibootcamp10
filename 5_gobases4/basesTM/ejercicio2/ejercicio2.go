package main

import (
	"errors"
	"fmt"
)

func verificarSalario(salary int) error {
	if salary < 150000 {
		return errors.New("el salario ingresado no alcanza el minimo imponible")
	}
	return nil
}

func main() {
	var salary int
	fmt.Println("Ingrese su salario")
	fmt.Scanln(&salary)
	e := verificarSalario(salary)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
