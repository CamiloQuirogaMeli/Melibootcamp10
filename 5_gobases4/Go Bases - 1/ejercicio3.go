package main

import (
	"fmt"
)

/*
	Ejercicio 3 - Impuestos de salario #3

	Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
	error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
	(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
	el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

func validaSalario(valor int) error {
	sueldoRequerido := 150000
	if valor < sueldoRequerido {
		err := fmt.Errorf("error: el mínimo imponible es de %d y el salario ingresado es de: %d", sueldoRequerido, valor)
		return err
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
