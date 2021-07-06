package main

import (
	"fmt"
	"math"
)

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos []Productos, c chan float64) {
	var total float64
	for _, producto := range productos {
		total += producto.Precio * float64(producto.Cantidad)
	}
	c <- total
}

func sumarServicios(servicios []Servicios, c chan float64) {
	var total float64
	for _, servicio := range servicios {
		if servicio.MinutosTrabajados < 30 {
			total += servicio.Precio
		} else {
			fraccionTiempo := math.Ceil(float64(servicio.MinutosTrabajados) / 30)
			total += servicio.Precio * fraccionTiempo
		}
	}
	c <- total
}

func sumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	var total float64
	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.Precio
	}
	c <- total
}

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	p1 := Productos{Nombre: "Alfajores", Precio: 5, Cantidad: 10}
	p2 := Productos{Nombre: "Milanesas", Precio: 4, Cantidad: 20}
	productos := []Productos{p1, p2}
	go sumarProductos(productos, c1)

	s1 := Servicios{Nombre: "Reparacion PC", MinutosTrabajados: 20, Precio: 100.00}
	s2 := Servicios{Nombre: "Reparacion Heladera", MinutosTrabajados: 60, Precio: 1000.00}
	servicios := []Servicios{s1, s2}
	go sumarServicios(servicios, c2)

	m1 := Mantenimiento{Nombre: "Plomeria", Precio: 250}
	m2 := Mantenimiento{Nombre: "Gasista", Precio: 1500}
	mantenimiento := []Mantenimiento{m1, m2}
	go sumarMantenimiento(mantenimiento, c3)

	total := <-c1 + <-c2 + <-c3
	fmt.Println(total)
}
