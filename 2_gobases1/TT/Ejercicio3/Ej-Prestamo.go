package main

import "fmt"

func main() {
	var edad uint16 = 23
	var antAños uint16 = 3
	var sueldo float32 = 100000

	if edad > 22 && antAños > 1 {

		if sueldo > 100000 {
			fmt.Println(("Se otorgará prestamo sin intereses."))
		} else {
			fmt.Println(("Se otorgará prestamo con intereses."))
		}

	} else {
		fmt.Println(("No se otorgará prestamo."))
	}
}
