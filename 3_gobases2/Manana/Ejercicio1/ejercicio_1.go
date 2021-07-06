package main

import "fmt"

func calcularSueldo(sueldo float64) float64 {
	if sueldo > 50000 {
		if sueldo > 150000 {
			return sueldo * 0.73
		} else {
			return sueldo * 0.83
		}
	} else {
		return sueldo
	}
}

func main() {
	sueldo := 45000.0
	fmt.Println("El sueldo neto es: ", calcularSueldo(sueldo))
}
