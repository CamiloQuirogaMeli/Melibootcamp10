package main

import "fmt"

func main() {

	/*
		Ejercicio 1 - Impuestos de salario #1
		1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
		“int”.
		2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
		“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
		“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
		pagar impuesto”.
	*/
	e := myError{"error: el salario ingresado no alcanza el mínimo imponible"}

	var salary int
	fmt.Println("Ingrese salario")
	fmt.Scanln(&salary)

	if salary < 150000 {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

type myError struct {
	msj string
}

func (e *myError) Error() string {

	return e.msj
}
