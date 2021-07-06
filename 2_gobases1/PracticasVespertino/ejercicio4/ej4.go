package main

import (
	"fmt"
)

var (
	month = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Setimebre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre"}
	opc int
)

func main() {

	fmt.Println("Ingrese numero de mes")
	fmt.Scanf("%d", &opc)

	if opc >= 1 && opc <= 12 {
		fmt.Println(month[opc])
	} else {
		println("opcion no valida")
	}

}

/*
	 Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
	que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
	por qué?
	Ej: 7, Julio
*/
