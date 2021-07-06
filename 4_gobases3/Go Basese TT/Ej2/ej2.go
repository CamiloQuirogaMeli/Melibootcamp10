package main

import (
	"fmt"
	"os"
)

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}
type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

func main() {
	opcion := 0
	for opcion != 4 {
		fmt.Println("1. Crear producto")
		fmt.Println("2. Agregar producto")
		fmt.Println("3. Eliminar productos")
		fmt.Println("4. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
		case 2:
		case 3:
		case 4:
			os.Exit(1)
		}
	}
}
