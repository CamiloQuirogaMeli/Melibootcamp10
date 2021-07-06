package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(productos []Producto, c chan float64) {
	var precioTotal float64
	for _, producto := range productos {
		precioTotal += producto.precio * float64(producto.cantidad)
	}
	fmt.Println("Precio total de los productos: ", precioTotal)
	c <- precioTotal
}

func sumarServicios(servicios []Servicio, c chan float64) {
	var precioTotal float64
	for _, servicio := range servicios {
		var precioServicio float64
		if servicio.minutosTrabajados%30 != 0 {
			precioServicio += float64((servicio.minutosTrabajados / 30) + 1)
		} else {
			precioServicio += float64(servicio.minutosTrabajados / 30)
		}
		precioServicio *= servicio.precio
		precioTotal += precioServicio
	}
	fmt.Println("Precio total de los servicios: ", precioTotal)
	c <- precioTotal
}

func sumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	var precioTotal float64
	for _, mantenimiento := range mantenimientos {
		precioTotal += mantenimiento.precio
	}
	fmt.Println("Precio total de los mantenimientos: ", precioTotal)
	c <- precioTotal
}

func main() {
	c := make(chan float64)
	var montoFinal float64
	var productos = []Producto{
		{"celular", 10000, 2},
		{"tostadora", 200, 1},
		{"medias", 100, 2},
	}
	var servicios = []Servicio{
		{"Limpieza", 100, 60},
		{"Cocina", 300, 33},
	}
	var mantenimientos = []Mantenimiento{
		{"Arreglo televisor", 500},
		{"Arreglo 2", 300},
		{"Otro Mantenimiento", 150},
	}

	go sumarProductos(productos, c)
	go sumarServicios(servicios, c)
	go sumarMantenimiento(mantenimientos, c)
	for i := 0; i < 3; i++ {
		montoFinal += <-c
	}
	fmt.Println("Monto final: ", montoFinal)
}
