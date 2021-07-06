package main

import "fmt"

func main() {
	var numeroMes int
	var seguir bool = true

	for seguir {
		fmt.Println("Ingresá el número de mes: ")
		fmt.Scanln(&numeroMes)
		if numeroMes < 1 || numeroMes > 12 {
			fmt.Println("Número no válido")
		} else {
			switch numeroMes {
			case 1:
				fmt.Println("Enero")
			case 2:
				fmt.Println("Febrero")
			case 3:
				fmt.Println("Marzo")
			case 4:
				fmt.Println("Abril")
			case 5:
				fmt.Println("Mayo")
			case 6:
				fmt.Println("Junio")
			case 7:
				fmt.Println("Julio")
			case 8:
				fmt.Println("Agosto")
			case 9:
				fmt.Println("Septiembre")
			case 10:
				fmt.Println("Octubre")
			case 11:
				fmt.Println("Noviembre")
			case 12:
				fmt.Println("Diciembre")
			}
		}

		fmt.Println("¿Querés probar otro mes? Ingresa true o false")
		fmt.Scanln(&seguir)

	}

}
