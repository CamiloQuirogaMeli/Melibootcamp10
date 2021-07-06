package main

import "fmt"

func impuesto() float64 {

	var imp float64

	fmt.Print("Ingrese su sueldo: ")
	var salary float64
	fmt.Scanln(&salary)

	if salary > 50000 {
		imp = salary * 0.17
	} else if salary > 150000 {
		imp = salary * 0.27
	}
	return imp

}
