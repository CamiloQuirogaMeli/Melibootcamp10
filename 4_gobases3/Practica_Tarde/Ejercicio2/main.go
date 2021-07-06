package main

import "fmt"

type Usuario struct {
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

func nuevoProducto(nombre string, precio float64) (p Producto) {
	p.nombre = nombre
	p.precio = precio
	return
}

func agregarProducto(u *Usuario, pr *Producto, cantidad int) {
	(*u).productos = append((*u).productos, *pr)
}

func borrarProducto(u *Usuario) {
	(*u).productos = nil
}

func main() {
	u1 := Usuario{nombre: "Valeria", apellido: "Macina", correo: "valeria.macina@mercadolibre.com"}
	p1 := nuevoProducto("Producto1", 500)

	punteroProd := &p1
	punteroUsu := &u1

	agregarProducto(punteroUsu, punteroProd, 10)
	fmt.Print(u1, "\n")
	borrarProducto(punteroUsu)
	fmt.Print(u1)
}
