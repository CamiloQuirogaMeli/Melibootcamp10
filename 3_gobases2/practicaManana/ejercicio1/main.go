package main

import "fmt"

const TOPE1, TOPE2 float32 = 50000.00, 150000.00

func main() {
	var sueldo float32
	fmt.Println("Ingresar sueldo del empleado:")
	fmt.Scanln(&sueldo)
	fmt.Println("Total de impuestos: ", calcularImpuestos(sueldo))

}

func calcularImpuestos(sueldo float32) float32 {
	var impuesto float32 = 0.0
	if sueldo < TOPE1 {

	} else {
		impuesto = (sueldo * 17) / 100
		if sueldo > TOPE2 {
			impuesto = impuesto + ((sueldo * 10) / 100)
		}
	}
	return impuesto
}
