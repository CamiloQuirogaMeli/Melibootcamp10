package main

import "fmt"

type Products struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type User struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Products
}

func main() {
	var p1 Products
	p := &p1

	newProduct(p, "Papa", 96.4)
	fmt.Println(p1)

	u1 := User{Nombre: "Joni", Apellido: "Krucheski", Correo: "jonathankrucheski@gmail.com"}
	u := &u1
	fmt.Println(u1)

	addProduct(u, p, 4)
	fmt.Println(*u)

	deleteProducts(u)
	fmt.Println(u1)
}

func newProduct(p *Products, name string, price float64) {
	p.Nombre = name
	p.Precio = price
}

func addProduct(u *User, p *Products, cant int) {
	p.Cantidad = cant
	u.Productos = append(u.Productos, *p)

}

func deleteProducts(u *User) {
	u.Productos = make([]Products, 0)
}
