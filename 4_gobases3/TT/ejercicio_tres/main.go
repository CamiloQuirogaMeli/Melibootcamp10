package main

import (
	"fmt"
)

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

type Servicios struct {
	Nombre            string
	Precio            float32
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float32
}

func sumarProducts(productos []Producto, c chan float32) (result float32) {

	for _, producto := range productos {
		result += producto.Precio * float32(producto.Cantidad)
	}
	c <- result
	return
}

func sumarServicios(servicios []Servicios, c chan float32) (result float32) {

	for _, servicio := range servicios {

		mediasHoras := 0
		if (servicio.MinutosTrabajados % 30) != 0 {
			mediasHoras++
		}
		mediasHoras += servicio.MinutosTrabajados

		result += float32(mediasHoras) * servicio.Precio
	}
	c <- result
	return
}

func sumarMantenimientos(mantenimientos []Mantenimiento, c chan float32) (result float32) {

	for _, mantenimiento := range mantenimientos {
		result += mantenimiento.Precio
	}
	c <- result
	return
}

func main() {

	c := make(chan float32)

	var productos = make([]Producto, 0)
	productos = append(productos, Producto{"n1", 456, 4})
	productos = append(productos, Producto{"n2", 300, 3})
	productos = append(productos, Producto{"n3", 250, 2})

	go sumarProducts(productos, c)

	var servicios = make([]Servicios, 0)
	servicios = append(servicios, Servicios{"s1", 3453, 80})
	servicios = append(servicios, Servicios{"s2", 344, 120})
	servicios = append(servicios, Servicios{"s3", 4543, 560})

	go sumarServicios(servicios, c)

	var mantenimientos = make([]Mantenimiento, 0)
	mantenimientos = append(mantenimientos, Mantenimiento{"m1", 345})
	mantenimientos = append(mantenimientos, Mantenimiento{"m2", 3445})
	mantenimientos = append(mantenimientos, Mantenimiento{"m3", 454})

	go sumarMantenimientos(mantenimientos, c)

	var monto_total float32
	for i := 0; i < 3; i++ {
		monto_total += <-c
	}
	fmt.Println("El costo total es: ", monto_total)

}
