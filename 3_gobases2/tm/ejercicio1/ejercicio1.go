package main

import (
	"fmt"
)

const IMP17 float32 = 0.17
const IMP27 float32 = 0.27
const SUELDO50 float32 = float32(50000)
const SUELDO150 float32 = float32(150000)

func main() {
	var sueldo float32

	fmt.Println("Ingrese el sueldo del empleado:")
	fmt.Scanln(&sueldo)
	fmt.Printf("El impuesto es %.2f\n", impuesto(sueldo))

}

func impuesto(sueldo float32) float32 {
	switch {
	case sueldo >= SUELDO50 && sueldo < SUELDO150:
		return sueldo * IMP17
	case sueldo >= SUELDO150:
		return sueldo * IMP27
	default:
		return float32(0)
	}
}
