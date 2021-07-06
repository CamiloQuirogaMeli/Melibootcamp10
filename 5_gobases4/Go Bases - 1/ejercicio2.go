package main

import (
	"errors"
	"fmt"
)

/*
	Ejercicio 2 - Impuestos de salario #2
	Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
	reemplazo de “Error()”, se implemente “errors.New()”.
*/

func validaSalario(valor int) error {
	if valor < 150000 {
		return errors.New("el salario ingresado no alcanza el mínimo imponible")
	}
	return nil

}

func main() {
	var salary int = 0
	fmt.Print("Introduzca el salario: ")
	fmt.Scan(&salary)
	e := validaSalario(salary)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
