package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Ejercicio 2 - Impuestos de salario #2
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”.

*/

func checkSalary(salary int) error {

	if salary < 150000 {
		return errors.New("el salario ingresado no alcanza el mínimo imponible")
	}

	return nil

}

func main() {
	salary := 150000

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
