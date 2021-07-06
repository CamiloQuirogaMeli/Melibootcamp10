// Ejercicio 3 - Impuestos de salario #3
// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary”
// indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”,
// siendo [salary] el valor de tipo int pasado por parámetro).

package main

import (
	f "fmt"
)

func verificateSalary(salary int) error {
	if salary < 150000 {
		return f.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: [%d]", salary)
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
