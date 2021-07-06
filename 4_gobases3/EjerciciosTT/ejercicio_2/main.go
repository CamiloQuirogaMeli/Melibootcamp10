package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

func nuevoProducto(nombre string, precio float64) *Producto {
	return &Producto{
		nombre: nombre,
		precio: precio,
	}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.cantidad = cantidad
	usuario.productos = append(usuario.productos, producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.productos = []*Producto{}
}

func main() {
	u := &Usuario{
		nombre:   "Mario",
		apellido: "Rosales",
		correo:   "marioalberto.rosales@mercadolibre.com",
	}

	producto := nuevoProducto("mouse", 124.34)
	producto2 := nuevoProducto("teclado", 564.49)
	agregarProducto(u, producto, 2)
	agregarProducto(u, producto2, 5)
	fmt.Println(u.nombre, u.apellido, u.correo)
	for _, valor := range u.productos {
		fmt.Println(*valor)
	}
}
