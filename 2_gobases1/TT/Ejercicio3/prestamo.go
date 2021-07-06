package main

import "fmt"

func main() {
	var edad int
	var sueldo float32
	var antiguedad int

	fmt.Print("edad: ")
	fmt.Scan(&edad)

	fmt.Print("sueldo: ")
	fmt.Scan(&sueldo)

	fmt.Print("antiguedad: ")
	fmt.Scan(&antiguedad)

	if edad > 22 && antiguedad > 1 {
		fmt.Println("Se asigna el credito")
		if sueldo > 100000 {
			fmt.Println("No se cobra interes")
		} else {
			fmt.Println("Se cobra interes")
		}
	} else if edad > 22 && antiguedad < 2 {
		fmt.Println("su antiguedad debe ser mayor a 1 año")
	} else if edad < 23 && antiguedad > 1 {
		fmt.Println("La edad debe ser mayor a 22")
	} else {
		fmt.Println("La edad debe ser mayor a 22 y la antiguedad mayor a 1 año")
	}
}
