package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{Nombre: nombre, Precio: precio}
}

func aregarProducto(u *Usuario, p *Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
}

func borrarProductos(u *Usuario) {
	u.Productos = []*Producto{}
}

func main() {
	usuario := &Usuario{Nombre: "Carlos", Apellido: "Perez", Correo: "test@gmail.com"}
	alfajor := nuevoProducto("Alfajor", 10.50)
	aregarProducto(usuario, alfajor, 2)
	empanadas := nuevoProducto("Empanadas", 12)
	aregarProducto(usuario, empanadas, 2)

	for _, producto := range usuario.Productos {
		fmt.Println(*producto)
	}

	borrarProductos(usuario)

	fmt.Println(usuario.Productos)
}
