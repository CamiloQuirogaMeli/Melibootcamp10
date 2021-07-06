package main

import (
	"fmt"
)

func ejercicio4() {

	fede := User{
		Nombre:   "Federico",
		Apellido: "Martinez",
		Correo:   "Fedemartinez@gmail.com",
	}

	fmt.Println(fede)

	fideos := newProduct("fideos", 50, 1)

	fmt.Println(fideos)

	agregarProducto(&fede, &fideos, 50)

	fmt.Println(fede)

	fmt.Println(fideos)

	borrarProducto(&fede)

	fmt.Println(fede)
}

type User struct {
	Nombre   string
	Apellido string
	Correo   string
	Product  []Producto
}

func agregarProducto(user *User, producto *Producto, cantidad float64) {

	producto.Cantidad = cantidad
	user.Product = append(user.Product, *producto)
}

func borrarProducto(user *User) {
	user.Product = nil
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad float64
}

func newProduct(nombre string, precio float64, cantidad float64) Producto {

	return Producto{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}
