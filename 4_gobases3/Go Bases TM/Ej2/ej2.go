package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Producto struct {
	Id       int
	Precio   float32
	Cantidad int
}

func main() {
	archivo, err := ioutil.ReadFile("../Ej1/archivoGenerado.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo")
		os.Exit(1)
	}
	fmt.Println("ID\tPrecio\tCantidad")
	fmt.Println(string(archivo))

}
