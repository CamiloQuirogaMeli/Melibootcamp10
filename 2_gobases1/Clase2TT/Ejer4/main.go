package main

import (
	"fmt"
)

func main() {
	var meses = map[int]string{
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
		12: "diciembre"}
	mes := 10
	if mes < 1 || mes > 12 {
		fmt.Println("El mes no existe")
	} else {
		fmt.Println(meses[mes])
	}
}
