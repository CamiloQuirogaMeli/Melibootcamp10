package main

import (
	"fmt"
	"io/ioutil"
)

type Producto struct {
	id       string
	precio   string
	cantidad string
}

func saveData(productos []Producto) {
	var msg string

	for i, p := range productos {
		msg += p.id + "," + p.precio + "," + p.cantidad
		if (i + 1) < len(productos) {
			msg += ";"
		}
	}
	data := []byte(msg)

	err := ioutil.WriteFile("/Users/mrosales/Desktop/bootcamp/meli_bootcamp10/4_gobases3/EjerciciosTM/ejercicio_1/productos.txt", data, 777)

	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
	}
}

func main() {
	p1 := Producto{id: "1", precio: "13450.98", cantidad: "57"}
	p2 := Producto{id: "2", precio: "43332.58", cantidad: "3456"}
	p3 := Producto{id: "3", precio: "1340", cantidad: "989876"}

	var productos []Producto

	productos = append(productos, p1)
	productos = append(productos, p2)
	productos = append(productos, p3)

	saveData(productos)
}
