// Ejercicio 2 - Ecommerce
// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
// tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.

package main

import (
	f "fmt"
	s "strings"
)

type User struct {
	name     string
	lastname string
	mail     string
	products []Product
}

type Product struct {
	name  string
	price float64
	cant  int
}

/// No se utiliza puntero, si retorna un producto no le veo sentido que sea con punteros.
func newProduct(nameProduct string, price float64) Product {
	var p Product

	p.name = nameProduct
	p.price = price

	return p
}

func deleteProducts(user *User) {
	user.products = user.products[:(len(user.products) - 1)]
}

func addProduct(user *User, prod *Product, cant int) {
	prod.cant = cant
	user.products = append(user.products, *(prod))
}

func showUserAndProducts(user User) {
	f.Printf("Nombre de usuario: %s\n", user.name)
	f.Printf("Apellido de usuario: %s\n", user.lastname)
	f.Printf("E-mail del usuario: %s\n", user.mail)

	f.Printf("Productos que pose el usuario:\n")

	for _, prod := range user.products {
		f.Printf("Nombre del producto: %s\n", prod.name)
		f.Printf("Precio del producto: %.2f\n", prod.price)
		f.Printf("Cantidad del producto: %d\n", prod.cant)
	}
	f.Println()
}

func main() {

	var nameProduct string
	var priceProduct float64
	var cantProduct int
	var user User
	var option string = "si"

	f.Printf("Ingrese el nombre de usuario al que desea agregar los productos: ")
	f.Scan(&user.name)
	f.Printf("Ingrese el apellidp del usuario al que desea agregar los productos: ")
	f.Scan(&user.lastname)
	f.Printf("Ingrese el mail del usuario al que desea agregar los productos: ")
	f.Scan(&user.mail)

	for s.Compare(option, "si") == 0 {
		f.Printf("Ingrese el nombre del producto que desea agregar: ")
		f.Scan(&nameProduct)
		f.Printf("Ingrese el precio del producto que desea agregar: ")
		f.Scan(&priceProduct)
		f.Printf("Ingrese la cantidad del producto que desea agregar: ")
		f.Scan(&cantProduct)
		product := newProduct(nameProduct, priceProduct)
		addProduct(&user, &product, cantProduct)
		f.Printf("¿Desea agregar otro producto?(si/no): ")
		f.Scan(&option)
		option = s.ToLower(option)
	}
	f.Println()

	showUserAndProducts(user)

	f.Println("Ahora se borraran los productos del usuario")

	deleteProducts(&user)

	f.Println()

	showUserAndProducts(user)
}
