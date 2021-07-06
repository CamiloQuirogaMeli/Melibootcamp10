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

func NuevoProducto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 0}
}

func AgregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.cantidad = cantidad
	usuario.productos = append(usuario.productos, *producto)
}

func BorrarProductos(usuario *Usuario) {
	usuario.productos = make([]Producto, 0)
}

func main() {
	usuario := Usuario{"Luis", "Oropeza", "luis@mercadolibre.com.mx", make([]Producto, 0)}
	fmt.Println("Usuario:", usuario)
	producto := NuevoProducto("p1", 10.0)
	AgregarProducto(&usuario, &producto, 10)
	fmt.Println("Usuario actualizado (1):", usuario)
	producto = NuevoProducto("p2", 20.0)
	AgregarProducto(&usuario, &producto, 20)
	fmt.Println("Usuario actualizado (2):", usuario)
	BorrarProductos(&usuario)
	fmt.Println("Usuario actualizado (3):", usuario)
}
