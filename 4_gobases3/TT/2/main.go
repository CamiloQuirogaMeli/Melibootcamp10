package main

import (
	"fmt"
)

func main() {

	us := usuario{
		nombre:   "Lucas",
		apellido: "Ficarra",
		correo:   "example@gmail.com",
	}

	fmt.Print(us)

	prod := Producto{
		nombre: "remera",
		precio: 500.00,
	}

	AgregarProducto(&us, &prod, 5)

	fmt.Print(us)

	BorrarProductos(&us)

	fmt.Print(us)
}

type usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(nombre string, precio float64) Producto {
	prod := Producto{
		nombre: nombre,
		precio: precio,
	}
	return prod
}

func AgregarProducto(u *usuario, p *Producto, c int) {
	p.cantidad = c
	u.productos = append(u.productos, *p)
}

func BorrarProductos(u *usuario) {
	u.productos = []Producto{}
}
