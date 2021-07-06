package main

import "fmt"

func main() {
	var (
		edad      = 17
		empleado  = true
		añosAntig = 2
		sueldo    = 145000
	)
	if edad > 22 && empleado && añosAntig > 1 {
		if sueldo > 100000 {
			fmt.Println("Puede recibir prestamo SIN interes.")
		} else {
			fmt.Println("Puede recibir prestamo CON interes.")
		}
	} else {
		fmt.Println("NO puede recibir prestamo")
	}
}
