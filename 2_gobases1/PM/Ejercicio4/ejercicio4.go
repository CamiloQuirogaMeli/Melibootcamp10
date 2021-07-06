package main

import "fmt"

func main() {
	var numeroMes int
	var seguir bool = true

	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril",
		5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre",
		11: "Noviembre", 12: "Diciembre"}

	for seguir {
		fmt.Println("Ingresá el número de mes: ")
		fmt.Scanln(&numeroMes)
		if numeroMes < 1 || numeroMes > 12 {
			fmt.Println("Número no válido")
		} else {
			fmt.Println(meses[numeroMes])
		}

		fmt.Println("¿Querés probar otro mes? Ingresa true o false")
		fmt.Scanln(&seguir)

	}

	for i, mes := range meses {
		fmt.Println(i, mes)

	}

}
