package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		Ejercicio 2 - Impuestos de salario #2
		Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
		reemplazo de “Error()”, se implemente “errors.New()”.
	*/

	e := errors.New("error: el salario ingresado no alcanza el mínimo imponible")

	var salary int
	fmt.Println("Ingrese salario")
	fmt.Scanln(&salary)

	if salary < 150000 {
		fmt.Println(e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
