package main

import "fmt"

type error interface {
	Error() string
}

type minSalaryError struct {
	status int
	msg string
}

func (e minSalaryError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func checkSalary (salary int) error {
	if salary < 150000 {
		return minSalaryError{}
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
