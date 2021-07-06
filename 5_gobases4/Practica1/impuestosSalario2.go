package main

import (
	"fmt"
	"errors"
)

//Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”, se implemente “errors.New()”.

func validSalary (salary int) (error) {

	if salary < 150000 {
	   	return errors.New("error: el salario ingresado no alcanza el mínimo imponibl")
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