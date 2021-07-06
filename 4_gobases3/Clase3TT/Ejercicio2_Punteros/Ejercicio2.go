package main

import (
	"fmt"
)

// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
// productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
// tengan la misma dirección de memoria en el main del programa como en las funciones.
// Se necesitan las estructuras:
// - Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// - Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// - Nuevo producto: recibe nombre y precio, y retorna un producto.
// - Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el
// producto al usuario.
// - Borrar productos: recibe un usuario, borra los productos del usuario.

type User struct {
	Name     string
	Lastname string
	Email    string
	Products []Product
}

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func newProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func addProduct(user *User, prod *Product, quant int) {
	prod.Amount = quant
	user.Products = append(user.Products, *prod)
}

func delProduct(user *User) {
	user.Products = user.Products[:0]
}

func main() {

	user := User{Name: "Martin", Lastname: "Hernandez", Email: "mail@gmail"}
	prod1 := newProduct("Milanesa", 400)
	prod2 := newProduct("Ensalada", 350)
	prod3 := newProduct("Pasta", 380)

	addProduct(&user, prod1, 3)
	addProduct(&user, prod2, 2)
	addProduct(&user, prod3, 5)

	fmt.Println(user)

}
