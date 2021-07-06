package main

import (
	"fmt"
)

func checkSalary (salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: [%d]", salary)
	}
	return nil
}

func main(){
	salary := 150000 - 1
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}
