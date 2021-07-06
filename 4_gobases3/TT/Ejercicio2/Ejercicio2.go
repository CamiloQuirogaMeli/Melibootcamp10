package main

import "fmt"

func main() {
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
	p1 := newProduct("Object", 10000)
	p2 := newProduct("Object2", 15000)

	u1 := User{}
	u1.name = "Name"
	u1.surname = "Surname"
	u1.email = "name@correo.com"

	addProduct(&u1, &p1, 3)
	addProduct(&u1, &p2, 3)
	fmt.Println(u1)
	deleteProduct(&u1)
	fmt.Println(u1)
}

type User struct {
	name    string
	surname string
	email   string
	product []Product
}

type Product struct {
	name  string
	price float64
	cant  int
}

func newProduct(newName string, newPrice float64) (newProd Product) {
	newProd.name = newName
	newProd.price = newPrice
	return
}

func addProduct(u *User, p *Product, cant int) {
	(*u).product = append((*u).product, *p)
	p.cant = cant
}

func deleteProduct(u *User) {
	(*u).product = []Product{}
}
