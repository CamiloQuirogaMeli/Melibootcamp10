package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("productos.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo")
		os.Exit(1)
	}
	datos := strings.Split(string(content), "\n")
	fmt.Printf("ID\tPrecio\tCantidad\n")
	for _, linea := range datos {
		if len(linea) == 0 {
			continue
		}
		producto := strings.ReplaceAll(linea, ";", "\t")
		fmt.Println(producto)
	}
}
