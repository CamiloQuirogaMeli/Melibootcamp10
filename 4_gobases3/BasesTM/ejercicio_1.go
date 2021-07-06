package main

import (
	"fmt"
	"io/ioutil"
)

type Producto struct {
	id       int
	precio   float32
	cantidad int
}

func (p Producto) String() string {
	return fmt.Sprintf("%d;%0.2f;%d\n", p.id, p.precio, p.cantidad)
}

func main() {
	productos := []Producto{{1, 10.50, 10}, {2, 1000.0, 2}}
	var listaProductos string
	for _, p := range productos {
		listaProductos += p.String()
	}
	info := []byte(listaProductos)
	err := ioutil.WriteFile("productos.txt", info, 0644)
	if err != nil {
		fmt.Println("Error al guardar la informacion")
	}
}
