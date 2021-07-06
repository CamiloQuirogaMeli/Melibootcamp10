package main

import "fmt"

func main() {
	/* Solucion con Array

	var months [12]string
	var selectedMonth int = 12

	months[0] = "Enero"
	months[1] = "Febrero"
	months[2] = "Marzo"
	months[3] = "Abril"
	months[4] = "Mayo"
	months[5] = "Junio"
	months[6] = "Julio"
	months[7] = "Agosto"
	months[8] = "Septiembre"
	months[9] = "Octubre"
	months[10] = "Noviembre"
	months[11] = "Diciembre"

	fmt.Println(months[selectedMonth-1]) */

	/*Solucion con switch
	var selectedMonth int = 8

	switch selectedMonth {
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
	default:
		fmt.Println("Ingrese un mes valido")
	}
	*/
	/* Solucion con Mapa - Esta solucion me resulta mas legible y entiendo que es la mas rapida*/

	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio",
		7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	var selectedMonth int
	fmt.Print("Ingrese un mes en numeros: ")
	fmt.Scanf("%d", &selectedMonth)
	if selectedMonth > 12 {
		fmt.Println("Ingrese un mes valido")
	} else {
		fmt.Println(months[selectedMonth])
	}
}
