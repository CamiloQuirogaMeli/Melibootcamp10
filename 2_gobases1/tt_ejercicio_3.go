package main

import "fmt"

type clienType struct {
	id         int
	edad       int
	trabaja    bool
	antiguedad int
	sueldo     float32
}

func main() {

	clientes := []clienType{
		{123, 21, true, 2, 100000},
		{453, 25, true, 3, 100000},
		{4574, 19, true, 1, 70000},
		{456, 15, false, 0, 0},
		{234, 25, true, 3, 150000},
	}

	for i, cliente := range clientes {
		var verificacion int = verificarCliente(clientes[i])

		switch verificacion {
		case 0:
			fmt.Printf("El cliente %d no accede a un credito \n", cliente.id)
		case 1:
			fmt.Printf("El cliente %d accede a un credito \n", cliente.id)
		case 2:
			fmt.Printf("El cliente %d accede a un credito sin intereses \n", cliente.id)
		}

	}
}

func verificarCliente(c clienType) int {
	if mayorDe22(c.edad) && antiguedadDe1(c.antiguedad) && c.trabaja && sueldoMayor100k(c.sueldo) {
		return 2
	}
	if mayorDe22(c.edad) && antiguedadDe1(c.antiguedad) && c.trabaja {
		return 1
	}
	return 0
}

func mayorDe22(edad int) bool {
	return edad > 22
}
func antiguedadDe1(antiguedad int) bool {
	return antiguedad > 1
}
func sueldoMayor100k(sueldo float32) bool {
	return sueldo > 100000
}
