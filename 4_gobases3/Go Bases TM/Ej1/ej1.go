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
	opcion := 0
	productos := []Producto{}
	for opcion != 3 {
		fmt.Println("1. Agregar Producto")
		fmt.Println("2. Generar archivo")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			prod := Producto{}
			fmt.Println("Identificador: ")
			fmt.Scanln(&prod.Id)
			fmt.Println("Precio: ")
			fmt.Scanln(&prod.Precio)
			fmt.Println("Cantidad: ")
			fmt.Scanln(&prod.Cantidad)
			productos = append(productos, prod)
		case 2:
			var escrito string
			for _, valor := range productos {
				escrito += fmt.Sprintf("%d\t%.2f\t%d\n", valor.Id, valor.Precio, valor.Cantidad)
			}
			bytes := []byte(escrito)
			ioutil.WriteFile("./archivoGenerado.txt", bytes, 0644)
			fmt.Println("Archivo Generado")
		case 3:
			os.Exit(1)
		}

	}
}
