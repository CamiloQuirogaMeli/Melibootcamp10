// Ejercicio 1 - Impuestos de salario #1
// 1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo “int”.
// 2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible"
// y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

package main

import (
	f "fmt"
)

type error interface {
	Error() string
}

type myCustomErrorSalary struct {
	msg string
}

func (e myCustomErrorSalary) Error() string {
	return "error: el salario ingresado no alcanza el minimo imponible"
}

func verificateSalary(salary int) error {
	if salary < 150000 {
		return myCustomErrorSalary{}
	}
	return nil
}

func main() {

	var salary int

	f.Printf("Ingrese el valor del salario: ")
	f.Scan(&salary)

	err := verificateSalary(salary)

	if err != nil {
		f.Println(err.Error())
	} else {
		f.Println("Debe pagar impuestos")
	}

}
