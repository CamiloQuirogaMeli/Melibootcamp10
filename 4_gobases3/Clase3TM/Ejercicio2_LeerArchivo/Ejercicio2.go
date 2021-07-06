// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
// por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
// y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
// visualizar el total (Sumando precio por cantidad)

// Ejemplo:
// ID 			Precio 		Cantidad
// 111223 		30012.00 	1
// 444321 		1000000.00 	4
// 434321 		50.50 		1

// 			4030062.50

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	dat, err := ioutil.ReadFile("./pruebaGuardado.txt")
	if err != nil {
		fmt.Printf("Algo salio mal al leer el archivo")
		fmt.Println(err)
	} else {
		fmt.Println("Id\t\t Precio\t \tCantidad\n")

		//Lo separo por lineas en un slice
		str_slice := strings.Split(string(dat), "\n")

		var total float64 = 0
		var column []string

		//Recorro el slice y lo vuelvo a separar por ;
		for i := 0; i < len(str_slice); i++ {
			column = strings.Split(str_slice[i], ";")
			// for i := 0; i < len(column); i++ {
			// 	println(column[i])
			// }
			col1, err := strconv.ParseFloat(column[1], 64)
			col2, err := strconv.ParseFloat(column[2], 64)
			if err != nil {
				fmt.Println(err)
			} else {
				total = total + (col1 * col2)
			}
			fmt.Println(column[0], "\t", column[1], "\t\t", column[2])
		}
		fmt.Println("\nTotal:", total)
		// for i := 0; i < len(str_slice); i++ {
		// 	fmt.Printf("%s \t", str_slice[i])
		// }
		// fmt.Printf("\n")
	}
}
