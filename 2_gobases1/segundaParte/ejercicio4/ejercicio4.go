package main

import "fmt"

func main() {
	var month int
	fmt.Println("Ingrese un numero de mes: ")
	fmt.Scanln(&month)
	var calendar = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	fmt.Printf("El mes es %s\n", calendar[month])

}
