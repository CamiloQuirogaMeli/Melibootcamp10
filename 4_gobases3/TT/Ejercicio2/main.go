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
	Precio   float64
	Cantidad int
}

func nuevoProducto(nom string, pre float64) (producto Producto) {
	producto.Nombre = nom
	producto.Precio = pre
	return
}

func agregarProducto(usu *Usuario, prod *Producto, cant int) {
	(*usu).Productos = append((*usu).Productos, *prod)
}

func borrarProductos(usu *Usuario) {
	(*usu).Productos = nil
}

func main() {
	u1 := Usuario{"Juan", "Perez", "pepito@hotmail.com", nil}
	p1 := nuevoProducto("Coca-cola", 152)
	fmt.Print(u1)

	puntUsu := &u1
	puntProd := &p1

	agregarProducto(puntUsu, puntProd, 10)
	fmt.Print(u1)

	borrarProductos(puntUsu)
	fmt.Print(u1)

}
