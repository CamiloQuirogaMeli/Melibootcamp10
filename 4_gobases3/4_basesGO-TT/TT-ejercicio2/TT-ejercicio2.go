package main

import "fmt"

/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
- Nuevo producto: recibe nombre y precio, y retorna un producto.
- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el
producto al usuario.
- Borrar productos: recibe un usuario, borra los productos del usuario.

*/

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

func nuevoProducto(nombre string, precio float64) producto {
	var p producto
	p.nombre = nombre
	p.precio = precio
	return p
}

func agregaProducto(u *usuario, p *producto, cant int) {
	(*p).cant = cant
	(*u).productos = append((*u).productos, (*p))
	(*p).cant = 0
}

func borrarProductos(u *usuario) {
	var p []producto
	(*u).productos = p
}

func main() {

	/*Declaro algunos productos*/
	p1 := nuevoProducto("Cuaderno", 100.0)
	p2 := nuevoProducto("Cartuchera", 170.0)
	p3 := nuevoProducto("Lapicera", 70)
	p4 := nuevoProducto("Carpeta", 359.27)
	p5 := nuevoProducto("Lapiz", 54.0)
	p6 := nuevoProducto("Pincel", 70.0)

	/*Declaro algunos usuarios*/
	var prod1 []producto
	var prod2 []producto
	u1 := usuario{nombre: "Nombre1", apellido: "Apellido1", correo: "us1@gmail.com", productos: prod1}
	u2 := usuario{nombre: "Nombre2", apellido: "Apellido2", correo: "us2@gmail.com", productos: prod2}

	/*Declaro dos punteros, uno para productos y usuarios*/
	var pUsuario *usuario
	var pProducto *producto

	/*Agrego productos al usuario 1 y los muestro*/
	pUsuario = &u1
	pProducto = &p1
	agregaProducto(pUsuario, pProducto, 1)
	pProducto = &p2
	agregaProducto(pUsuario, pProducto, 2)
	pProducto = &p3
	agregaProducto(pUsuario, pProducto, 10)
	pProducto = &p4
	agregaProducto(pUsuario, pProducto, 3)

	fmt.Println("Productos del usuario 1")
	fmt.Println(u1.productos)

	/*Agrego productos al usuario 2 y los muestro*/
	pUsuario = &u2
	pProducto = &p5
	agregaProducto(pUsuario, pProducto, 1)
	pProducto = &p6
	agregaProducto(pUsuario, pProducto, 2)

	fmt.Println("Productos del usuario 2")
	fmt.Println(u2.productos)

	/*Borro los productos del usuario 2*/
	borrarProductos(pUsuario)

	/*Muestro nuevamente usuarios y productos actualizados*/
	fmt.Println("Productos actualizados del usuario 1")
	fmt.Println(u1.productos)

	fmt.Println("Productos actualizados del usuario 2")
	fmt.Println(u2.productos)

}
