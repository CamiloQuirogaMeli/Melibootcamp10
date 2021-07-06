package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var total float64 = 0
	txt, _ := ioutil.ReadFile("../Ejercicio1/archivoEj1")
	productos := strings.Split(string(txt), ";")
	fmt.Printf("ID\tPrecio\tCantidad\n")
	for _, values := range productos {
		dato := strings.Split(values, " ")
		for _, value := range dato {
			if value != "" {
				precio, _ := strconv.ParseFloat(dato[1], 64)
				cantidad, _ := strconv.ParseFloat(dato[2], 64)
				total += precio * cantidad
				fmt.Printf("%v\t", value)
			}
		}
		fmt.Printf("\n")
		fmt.Println("\t", total)
	}
}
