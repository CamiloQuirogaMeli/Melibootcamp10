package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func formatTextToProd(text string) producto {
	s := strings.Fields(text)
	id, _ := strconv.Atoi(s[0])
	precio, _ := strconv.ParseFloat(s[1], 64)
	cant, _ := strconv.Atoi(s[2])
	p := producto{id, precio, cant}
	return p

}

func showProdsDetails(prods []producto) {
	total := 0.0
	fmt.Println("ID\tPRECIO CANTIDAD")
	for _, prod := range prods {
		total += prod.precio * float64(prod.cantidad)
		fmt.Printf("%d\t%.2f\t%d\n", prod.id, prod.precio, prod.cantidad)
	}
	fmt.Println("\n\tTOTAL", total)
}

func main() {
	data, err := ioutil.ReadFile("../Ejercicio1/productos.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo")
	} else {
		s := strings.Split(string(data), ";")
		p := []producto{}
		for i := 0; i < len(s)-1; i++ {
			prod := formatTextToProd(s[i])
			p = append(p, prod)
		}
		showProdsDetails(p)
	}
}
