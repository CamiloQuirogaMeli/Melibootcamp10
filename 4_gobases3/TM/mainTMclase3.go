package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Print("Ingrese ejercicio a resolver (1 o 2):\n")
	var ejercicioAresolver int
	fmt.Scanln(&ejercicioAresolver)

	switch ejercicioAresolver {
	case 1:
		p1 := nuevoProducto("macbook", 1, 450000, 2)
		p2 := nuevoProducto("bb-8", 8, 30000, 9)
		p3 := nuevoProducto("mouse Redragon", 20, 5300, 20)

		productos := []Producto{p1, p2, p3}
		guardarArchivo(productos)
	case 2:
		dat, err := ioutil.ReadFile("../TM/lista_productos.txt")
		if err != nil {
			fmt.Println(err)
		} else {
			plainText := string(dat)
			products := formatearData(plainText)
			printFormatData(products)
		}
	}
}
