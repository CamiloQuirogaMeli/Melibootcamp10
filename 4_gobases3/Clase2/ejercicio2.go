package main

import (
	"fmt"
)

type Usuario struct {
	nombre    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(nombre string, precio float64) Producto {
	producto := Producto{nombre: nombre, precio: precio}

	fmt.Println("Producto creado exitosamente")
	return producto
}

func AgregarProductoAUsuario(u *Usuario, p *Producto, cantidad int) {
	p.cantidad = cantidad
	u.productos = append(u.productos, *p)
	fmt.Println("Productos:", cantidad, p.nombre, "agregado")

}

func BorrarProductosUsuario(u *Usuario) {
	u.productos = u.productos[:0]
	fmt.Println("Productos borrados")
}

func main() {
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
	var nombreProducto string
	var precioProducto float64
	var cantidadProductos int
	var opcionYN string

	usuario1 := Usuario{nombre: "Julián"}

	productos := make([]Producto, 0)

	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")

		fmt.Println("\t 1: Crear un nuevo producto")
		fmt.Println("\t 2: Ver productos disponibles")
		fmt.Println("\t 3: Agregar producto a un usuario")
		fmt.Println("\t 4: Borrar Productos de un usuario")
		fmt.Println("\t 5: Imprimir Datos del usuario")
		fmt.Println("\t 6: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el nombre del producto:")
			fmt.Scanln(&nombreProducto)
			fmt.Println("Digita el precio del producto:")
			fmt.Scanln(&precioProducto)
			productos = append(productos, NuevoProducto(nombreProducto, precioProducto))
		case 2:
			fmt.Println("Los productos disponibles son:")
			for i, pro := range productos {
				fmt.Println("Producto", i+1, ":")
				fmt.Println("\tNombre:", pro.nombre)
				fmt.Println("\tPrecio:", pro.precio)

			}
		case 3:
			fmt.Println("Digita el nombre del producto para agregar:")
			fmt.Scanln(&nombreProducto)
			for _, pro := range productos {
				if pro.nombre == nombreProducto {
					fmt.Println("Digita la cantidad")
					fmt.Scanln(&cantidadProductos)
					AgregarProductoAUsuario(&usuario1, &pro, cantidadProductos)
					break
				} else {
					fmt.Println("El producto no existe")
				}
			}
		case 4:
			fmt.Println("Estas seguro de borrar los productos (y/n):")
			fmt.Scanln(&opcionYN)
			if opcionYN == "y" || opcionYN == "Y" || opcionYN == "n" || opcionYN == "N" {

				if opcionYN == "y" || opcionYN == "Y" {
					BorrarProductosUsuario(&usuario1)
					fmt.Println("Datos borrados")
				} else {
					fmt.Println("Los datos no se borraron")
				}
			} else {
				fmt.Println("Opción no válida")
			}
		case 5:
			fmt.Println("Los datos del usuario son los siguientes:")
			fmt.Println("Nombre:", usuario1.nombre)
			for i, p := range usuario1.productos {
				fmt.Println("Producto", i+1, ":")
				fmt.Println("\tNombre", p.nombre)
				fmt.Println("\tPrecio", p.precio)
				fmt.Println("\tCantidad", p.cantidad)

			}

		case 6:
			salir = true
		}
	}
}
