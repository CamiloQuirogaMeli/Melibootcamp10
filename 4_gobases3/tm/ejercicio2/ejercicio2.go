package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	txt, err := ioutil.ReadFile("/Users/kcuadrado/meli_bootcamp10/meli_bootcamp10/4_gobases3/tm/ejercicio1.txt")

	if err != nil {
		fmt.Println("ERROR! No se puede leer el archivo")
	}

	products := strings.Split(string(txt), ";")

	fmt.Printf("%-10s %20s %10s\n", "ID", "Precio", "Cantidad")

	var sum float64

	for _, p := range products {
		data := strings.Split(p, ",")
		price, _ := strconv.ParseFloat(data[1], 64)
		cant, _ := strconv.ParseFloat(data[2], 64)
		sum += price * cant
		fmt.Printf("%-10s %20.2f %10.2f \n", data[0], price, cant)
	}

	fmt.Printf("%-10s %20.2f\n", " ", sum)
}
