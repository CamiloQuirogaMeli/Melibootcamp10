package main

import (
	"fmt"
)

func main() {
	var numMonth int
	fmt.Println("Por favor inserte el numero del mes a continuación")
	fmt.Scan(&numMonth)
	numberOfMonthReNew(numMonth)
}

// func numberOfMonth(numMonth int) {
// 	switch numMonth {
// 	case 1:
// 		fmt.Println("Enero")
// 	case 2:
// 		fmt.Println("Febrero")
// 	case 3:
// 		fmt.Println("Marzo")
// 	case 4:
// 		fmt.Println("Abril")
// 	case 5:
// 		fmt.Println("Mayo")
// 	case 6:
// 		fmt.Println("Junio")
// 	case 7:
// 		fmt.Println("Julio")
// 	case 8:
// 		fmt.Println("Agosto")
// 	case 9:
// 		fmt.Println("Septiembre")
// 	case 10:
// 		fmt.Println("Octubre")
// 	case 11:
// 		fmt.Println("Noviembre")
// 	case 12:
// 		fmt.Println("Diciembre")
// 	default:
// 		fmt.Println("No existe ese numero de mes")
// 	}
// }

// func numberOfMonthNew(numMonth int) {
// 	var months [12]string
// 	months[0] = "Enero"
// 	months[1] = "Febrero"
// 	months[2] = "Marzo"
// 	months[3] = "Abril"
// 	months[4] = "Mayo"
// 	months[5] = "Junio"
// 	months[6] = "Julio"
// 	months[7] = "Agosto"
// 	months[8] = "Septiembre"
// 	months[9] = "Octubre"
// 	months[10] = "Noviembre"
// 	months[11] = "Diciembre"
// if numMonth >= 12 || numMonth < 0 {
// 	fmt.Println("No existe ese numero de mes")
// } else {
// 	fmt.Println(months[numMonth-1])
// }
// }

func numberOfMonthReNew(numMonth int) {
	months := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if numMonth >= 12 || numMonth < 0 {
		fmt.Println("No existe ese numero de mes")
	} else {
		fmt.Println(months[numMonth-1])
	}
}
