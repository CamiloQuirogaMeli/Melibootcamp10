package main

import "fmt"

func main() {

	prestamo(24, true, 3, 10000000)

}

func prestamo(edad int, empleado bool, antiguedad int, sueldo float64) {

	if edad > 22 {
		if empleado == true {
			if(antiguedad > 1){
				if(sueldo > 100000){
					fmt.Println("El cliente deberá abonar interés en su préstamo")
				}

			}else{
				fmt.Println("El cliente no lleva más de un año de antiguedad en su trabajo")
			}
		} else {
			fmt.Println("El cliente no se encuentra empleado")
		}
	} else {
		fmt.Println("El cliente es menor de 22")
	}

}
