package main

import (
	"fmt"
)

func main() {
	var antiguedad, edad int = 2, 23
	var empleado = true
	var sueldo float64 = 1234.56
	if antiguedad > 1 && edad > 22 && empleado {
		if sueldo < 100000.0 {
			fmt.Println("prestamo aprovado")
		} else {
			fmt.Println("prestamo aprovado y ademas no se te cobraran intereses")
		}
	} else {
		fmt.Println("ups no es posible aprovar su prestamo")
	}

}
