package main

import "fmt"

func main() {
	var salary float64
	var tax float64
	fmt.Println("Ingrese el sueldo:")
	fmt.Scanln(&salary)

	switch {
	case salary > 150000:
		tax += (salary * 0.1)
		fallthrough
	case salary > 50000:
		tax += (salary * 0.17)
	default:
		tax = 0
	}

	fmt.Println("El impuesto del salario ingresado es: ", tax)
}
