package main

import "fmt"

func main() {

	var sueldoBruto float64

	fmt.Print("Ingrese el sueldo bruto del empleado: \n")
	fmt.Scanf("%f", &sueldoBruto)

	imp, sueNeto := calcImp(sueldoBruto)

	fmt.Printf("El empleado cobrará $ %v en concepto de sueldo. Se le retuvo $ %v en concepto de impuestos \n", imp, sueNeto)

}

func calcImp(sueldoBruto float64) (impuestos float64, sueldoNeto float64) {

	if sueldoBruto > 50000.0 && sueldoBruto < 150000.0 {
		impuestos = sueldoBruto * 0.83
		sueldoNeto = sueldoBruto - impuestos

	} else if sueldoBruto >= 150000.0 {

		impuestos = sueldoBruto * 0.73
		sueldoNeto = sueldoBruto - impuestos

	} else {
		impuestos = 0.0
		sueldoNeto = sueldoBruto
	}
	return impuestos, sueldoNeto

}
