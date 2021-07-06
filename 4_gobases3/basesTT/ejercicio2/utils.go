package main

type Usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []Producto
}

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}
