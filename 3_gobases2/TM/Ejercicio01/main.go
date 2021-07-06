package main

import "fmt"

func obtenerImpuesto(sueldo int) int {
	var impuesto int
	if sueldo >= 150000 {
		impuesto = (sueldo * 27) / 100
	} else if sueldo >= 50000 {
		impuesto = (sueldo * 17) / 100
	}
	return impuesto
}

func main() {
	impuesto := obtenerImpuesto(150000)
	fmt.Println("El impuesto es de:", impuesto)
}
