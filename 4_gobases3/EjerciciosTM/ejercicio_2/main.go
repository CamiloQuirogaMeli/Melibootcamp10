package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("/Users/mrosales/Desktop/bootcamp/meli_bootcamp10/4_gobases3/EjerciciosTM/ejercicio_2/productos.txt")

	if err != nil {
		log.Fatal(err)
	}

	productos := strings.Split(string(data), ";")

	fmt.Println("ID \t\t\t Precio \t\t Cantidad")

	var total float64

	for _, producto := range productos {
		datos := strings.Split(producto, ",")
		precio, _ := strconv.ParseFloat(datos[1], 64)
		cantidad, _ := strconv.ParseFloat(datos[2], 64)
		total += precio * cantidad
		fmt.Println(datos[0], "\t\t\t", datos[1], "\t\t", datos[2])
	}

	fmt.Println("\t\t\t", total)
}
