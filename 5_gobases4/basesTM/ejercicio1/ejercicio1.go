package main

import "fmt"

type miError struct {
	mensaje string
}

func (e *miError) Error() string {
	return fmt.Sprintf("%v", e.mensaje)
}

func verificarSalario(salary int) error {
	if salary < 150000 {
		return &miError{
			mensaje: "el salario ingresado no alcanza el minimo imponible",
		}
	}
	return nil
}

func main() {
	var salary int
	fmt.Println("Ingrese su salario")
	fmt.Scanln(&salary)
	e := verificarSalario(salary)
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
