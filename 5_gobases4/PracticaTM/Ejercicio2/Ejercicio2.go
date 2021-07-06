// Ejercicio 2 - Impuestos de salario #2
// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”, se implemente “errors.New()”.

package main

import (
	e "errors"
	f "fmt"
)

func verificateSalary(salary int) error {
	if salary < 150000 {
		return e.New("error: el salario ingresado no alcanza el minimo imponible")
	}
	return nil
}

func main() {

	var salary int

	f.Printf("Ingrese el valor del salario: ")
	f.Scan(&salary)

	err := verificateSalary(salary)

	if err != nil {
		f.Println(err)
	} else {
		f.Println("Debe pagar impuestos")
	}

}
