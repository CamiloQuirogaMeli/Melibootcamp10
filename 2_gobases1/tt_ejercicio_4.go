package main

import "fmt"

func getMonth(mes int) string {
	mesArray := [12]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}

	return mesArray[mes-1]
}

func main() {
	mes := 1
	fmt.Printf("Es el mes de %s \n", getMonth(mes))
}
