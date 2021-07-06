package ej1

import (
	"fmt"
)

type errorSalary struct {
	msg string
}

func (e *errorSalary) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func testSalario(salary int) error {
	if salary < 150000 {
		return &errorSalary{
			msg: "myError: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func ValidarSalario(salary int) {
	err := testSalario(salary)

	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	} else {
		fmt.Printf("Debe pagar impuesto")
	}

}
