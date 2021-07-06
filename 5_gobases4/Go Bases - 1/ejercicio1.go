package main

import (
	"fmt"
)

/*
	Ejercicio 1 - Impuestos de salario #1
	1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
	“int”.
	2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
	“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
	“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
	pagar impuesto”.
*/

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func validaSalario(valor int) error {
	if valor < 150000 {
		return &myError{
			msg: "el salario ingresado no alcanza el mínimo imponible",
		}
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
