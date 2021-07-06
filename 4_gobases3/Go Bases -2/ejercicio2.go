package main

import (
	"errors"
	"fmt"
)

/*
	Ejercicio 2 - Ecommerce
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

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   int
	cantidad int
}

func main() {
	usuario := &Usuario{
		nombre:   "ander",
		apellido: "torres",
		correo:   "ander@gmail.com",
	}
	var (
		nombre   string
		precio   int
		cantidad int
	)

	var num = 0

	fmt.Println("Ingrese el nombre del producto:")
	fmt.Scan(&nombre)
	fmt.Println("Ingrese el precio del producto:")
	fmt.Scan(&precio)
	producto := nuevoProducto(nombre, precio)
	//fmt.Println(producto)

	fmt.Println("Ingrese la cantidad del producto:")
	fmt.Scan(&cantidad)
	//fmt.Println(usuario)

	err := agregarProducto(usuario, producto, cantidad)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Desea eliminar productos del usuario: 0.NO  1.SI - RESPUESTA --> ")
		fmt.Scan(&num)
		if num == 0 {
			fmt.Println(usuario)
		} else if num == 1 {
			err := borrarProductos(usuario)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(usuario)
			}
		} else {
			fmt.Println("La opcion ingresada no corresponde")
		}
	}

}

//Recibe el nombre y precio del producto y devulve un nuevo producto
func nuevoProducto(nombre string, precio int) *Producto {
	p := &Producto{
		nombre: nombre,
		precio: precio,
	}
	return p
}

// Recibe un usuario y usa la funcion enviarProducto para añadir un nuevo producto
func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("No puede ingresar menos de 1 unidad")
	}
	producto.cantidad = cantidad
	usuario.enviarProducto(producto)
	return nil
}

//Recibe un producto y añade al slice de productos del usuario, el producto
func (usuario *Usuario) enviarProducto(producto *Producto) {
	usuario.productos = append(usuario.productos, *producto)
}

//Recibe un usuario valida si no tiene productos, si los tiene elimina el slice
func borrarProductos(usuario *Usuario) error {
	if usuario.productos == nil {
		msj := errors.New("El usuario no cuenta con productos registrados")
		return msj
	}
	usuario.productos = nil
	return nil
}
