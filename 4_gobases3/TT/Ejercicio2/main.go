package main

import "fmt"

func main() {
	u := usuario{
		Nombre:   "Luciano",
		Apellido: "Dominino",
		Correo:   "luciano.dominino@mercadolibre.com",
	}

	p := producto{
		Nombre:   "CocaCola",
		Precio:   100.00,
		Cantidad: 2,
	}

	p.nuevoProducto("CocaCola", 110.00)
	u.agregarProducto(u, p)
	fmt.Println("Usuario con su producto", u)

	u.eliminarProductos(u)
	fmt.Println("Usuario sin productos", u)

}

type usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	Compras  []producto
}

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p *producto) nuevoProducto(nombre string, precio float64) producto {
	p.Nombre = nombre
	p.Precio = precio
	return *p
}

func (u *usuario) agregarProducto(usuario usuario, producto producto) {
	u.Compras = append(u.Compras, producto)
}

func (u *usuario) eliminarProductos(usuario usuario) {
	u.Compras = nil
}
