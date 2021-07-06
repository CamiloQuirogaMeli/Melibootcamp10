package main

import (
	"fmt"
	"os"
)

type MiError struct {
}

func (e *MiError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	salary := 10000
	respuesta, err := pagarImpuesto(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(respuesta)
}
func pagarImpuesto(valor int) (string, error) {
	if valor < 150000 {
		return "", &MiError{}
	} else {
		return "Debe pagar impuesto", nil
	}
}
