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

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{Nombre: nombre, Precio: precio, Cantidad: 0}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	(*producto).Cantidad = cantidad
	(*usuario).Productos = append((*usuario).Productos, *producto)
}

func borrarProductos(usuario *Usuario) {
	(*usuario).Productos = nil
}

func (u Usuario) mostrar() {
	fmt.Printf("\nNombre: %s\n", u.Nombre)
	fmt.Printf("Apellido: %s\n", u.Apellido)
	fmt.Printf("Correo: %s\n", u.Correo)
	for i, producto := range u.Productos {
		fmt.Printf("\n** Producto %d **\n", i+1)
		fmt.Printf("Nombre: %s\n", producto.Nombre)
		fmt.Printf("Precio: %.2f\n", producto.Precio)
		fmt.Printf("Cantidad: %d\n", producto.Cantidad)
	}
	fmt.Printf("\n")
}

func main() {
	var productos []Producto
	usuario := Usuario{
		Nombre:    "Nina",
		Apellido:  "Samartin",
		Correo:    "nina@gmail.com",
		Productos: productos,
	}
	productoA := nuevoProducto("A", 500)
	productoB := nuevoProducto("B", 1500)
	agregarProducto(&usuario, &productoA, 10)
	agregarProducto(&usuario, &productoB, 4)
	usuario.mostrar()

	borrarProductos(&usuario)
	productoC := nuevoProducto("C", 475)
	agregarProducto(&usuario, &productoC, 7)
	usuario.mostrar()
}
