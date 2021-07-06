package main

import "fmt"

func main() {
	var sueldo float32
	fmt.Println("Ingrese su sueldo")
	fmt.Scanln(&sueldo)
	impuesto, porcentaje := calcularImpuesto(sueldo)
	fmt.Println("El impuesto es: ", impuesto)
	fmt.Println("El porcentaje es: ", porcentaje)
	fmt.Println("Sueldo a cobrar: ", sueldo-impuesto)
}

func calcularImpuesto(sueldo float32) (float32, float32) {
	if sueldo > 50000.0 && sueldo <= 150000.0 {
		return sueldo * .17, 17
	} else if sueldo > 150000.0 {
		return sueldo * .1, 10
	} else {
		return 0.0, 0.0
	}
}
