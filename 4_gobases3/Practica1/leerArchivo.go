package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

//La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda
//para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

func main() {
	dat, err := ioutil.ReadFile("productos.csv")

    if err != nil {
        fmt.Println("ERROR:", err)
    } else {
		dataS := strings.Split(string(dat), "\n")
		
		fmt.Println("ID\tPrecio\tCantidad")

		var line []string
		var precio float64
		precio = 0

		for i := 1; i < len(dataS); i++ {
			line = strings.Split(string(dataS[i]), ";")
			i2, err := strconv.ParseFloat(line[1], 64)
			i1, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				fmt.Println("ERROR:", err)
			} else {
				precio += i2 * i1
			}
			fmt.Println(line[0], "\t", line[1], "\t", line[2])
		}

		fmt.Println("\t", precio, "\t")
	}
}