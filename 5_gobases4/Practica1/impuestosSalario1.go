package main

import (
	"fmt"
)
//1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo “int”.
//2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
//“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

type error interface {
	Error() string
}
 
type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomError (salary int) error {

	if salary < 150000 {
	   	return &myError {
		  	msg: "error: el salario ingresado no alcanza el mínimo imponible",
	   	}
	}

	return nil
}

func main() {
	salary := 140000

	err := myCustomError(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}