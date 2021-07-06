package main

import (
	"fmt"
)

func main() {
	var month uint
	var year [12]string
	year[0] = "enero"
	year[1] = "febrero"
	year[2] = "marzo"
	year[3] = "abril"
	year[4] = "mayo"
	year[5] = "junio"
	year[6] = "julio"
	year[7] = "agosto"
	year[8] = "septiembre"
	year[9] = "octubre"
	year[10] = "noviembre"
	year[11] = "diciembre"
	fmt.Print("Ingrese el número del mes: ")
	fmt.Scanln(&month)
	if month <= 12 {
		fmt.Println("El mes número", month, "del año es", year[month-1])
	} else {
		fmt.Println("Número inválido")
	}
}
