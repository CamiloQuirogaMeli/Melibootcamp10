package main

import "fmt"

type miError struct {
	msg string
}

func (e *miError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func errorSalario(salario float32) error {
	if salario < 150000 {
		return &miError{"error: el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return nil
	}
}

func main() {
	var salary float32

	fmt.Printf("Ingrese su salario: ")
	fmt.Scanf("%f", &salary)

	err := errorSalario(salary)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
