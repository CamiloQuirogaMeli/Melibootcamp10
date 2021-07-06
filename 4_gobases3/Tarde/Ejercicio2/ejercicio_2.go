package main

import (
	"fmt"
)

type usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []producto
}

type producto struct {
	nombre string
	precio float64
	cant   int
}

func nuevoProducto(n string, p float64) *producto {
	p1 := producto{nombre: n, precio: p, cant: 1}
	return &p1
}

func agregarProducto(u *usuario, p *producto, c int) {
	p.cant = c
	u.productos = append(u.productos, *p)
}

func borrarProductos(u *usuario) {
	u.productos = []producto{}
}

func main() {
	u1 := usuario{}
	u1.nombre = "Adrian"
	u1.apellido = "Nanin"
	u1.correo = "adrian.nanin@mercadolibre.com"

	p1 := nuevoProducto("Computadora", 50000.0)
	fmt.Println(*p1)

	agregarProducto(&u1, p1, 5)
	fmt.Println(u1)

	borrarProductos(&u1)
	fmt.Println(u1)
}
