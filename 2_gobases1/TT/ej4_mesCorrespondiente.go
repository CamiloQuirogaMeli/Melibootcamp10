package main

import (
	"fmt"
)

func mesCorr() {
	var month int
	var monthsMap = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto",
		9: "Septiembre", 10: "Ocutbre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Println("Ingrese mes: ")
	fmt.Scanln(&month)

	fmt.Println("Mes: ", month, "es --> ", monthsMap[month])
}
