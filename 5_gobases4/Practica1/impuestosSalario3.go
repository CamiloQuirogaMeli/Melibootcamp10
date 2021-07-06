package main

import (
	"fmt"
)

//Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el
//mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el
//valor de tipo int pasado por parámetro).

func validSalary (salary int) (error) {

	if salary < 150000 {
	   	return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}

	return nil

}

func main() {

	salary := 140000

	err := validSalary(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}