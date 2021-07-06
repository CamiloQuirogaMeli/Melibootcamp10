package main

import (
	"errors"
	"fmt"
)

func checkSalary (salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
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
