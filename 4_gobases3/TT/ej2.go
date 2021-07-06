package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func (u *Usuario) agregarProducto(p Producto) {
	u.Productos = append(u.Productos, p)
}

func (u *Usuario) nuevoProducto(n string, p int) Producto {
	return Producto{n, p, 0}
}

func (u *Usuario) borrarProductos() {
	u.Productos = nil
}

func main() {
	u := Usuario{"Daniel", "Sanchez", "dani@gmail.com", nil}
	fmt.Println("inicializo usuario: ", u)

	p := Producto{"Dulce", 20, 3}
	fmt.Println("inicializo producto: ", p)

	u.agregarProducto(p)
	fmt.Println("Agrego producto al usuario: ", u)

	newP := Producto{}
	newP = u.nuevoProducto("Botella", 15)
	u.Productos = append(u.Productos, newP)
	fmt.Println("Nuevo producto al usuario: ", u)

	u.borrarProductos()
	fmt.Println("Elimino productos del usuario: ", u)
}
