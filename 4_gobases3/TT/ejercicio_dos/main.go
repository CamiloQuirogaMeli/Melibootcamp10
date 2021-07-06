package main

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

func nuevoProducto(nombre string, precio float32) Producto {

	return Producto{nombre, precio, 0}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {

	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

func main() {

	productos := []Producto{}
	usuario := Usuario{"Patricio", "gomez", "ej@ej.com", productos}

	productoToBeAddeded := nuevoProducto("Producto1", 345)

	agregarProducto(&usuario, &productoToBeAddeded, 10)
	borrarProductos(&usuario)

}
