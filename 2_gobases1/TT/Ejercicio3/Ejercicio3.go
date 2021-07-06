package main

import "fmt"

func main() {
	var (
		age        int     = 22
		isEmployed bool    = false
		experience float32 = 1.4
		salary     int     = 65000
	)

	if isEmployed && age >= 22 && experience > float32(1) {
		if salary > 100000 {
			fmt.Println("Esta persona recibira un prestamo sin interes")
		} else {
			fmt.Println("Esta persona recibira un prestamo con interes")
		}
	} else {
		fmt.Println("Esta persona no puede recibir prestamos")
	}

}
