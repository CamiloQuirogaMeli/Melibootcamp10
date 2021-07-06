/* Ejercicio 2 - Ecommerce
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
- Borrar productos: recibe un usuario, borra los productos del usuario. */
package main

import (
	"fmt"
)

type Producto struct {
	Nombre string
	precio float64
	cantidad int
}

type Usuario struct {
	Nombre string
	Apellido string
	Correo string
	Productos []Producto
}
func borrarProducto(user *Usuario) {
	user.Productos = nil
}

func agregarProducto(user *Usuario, producto Producto, cantidad int) {
	producto.cantidad = cantidad
	user.Productos = append(user.Productos, producto)
}

func nuevoProducto(nombre string, precio float64) Producto {
	var producto Producto = Producto{nombre, precio, 0}
	return producto
}

var Usuarios []Usuario
//var Productos []Producto

func main() {
	//Productos = append(Productos, nuevoProducto("Test", 100.0))
	Usuarios = append(Usuarios, Usuario{"Gonzalo", "Rodriguez", "email", nil})
	agregarProducto(&Usuarios[0], nuevoProducto("Test", 100.0), 3)

	//fmt.Println(Productos)
	fmt.Println(Usuarios)
	borrarProducto(&Usuarios[0])
	fmt.Println(Usuarios)
}