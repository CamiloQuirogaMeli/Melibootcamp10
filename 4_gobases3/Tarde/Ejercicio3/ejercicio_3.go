package main

import (
	"fmt"
)

type Producto struct {
	nombre string
	precio float64
	cant   uint
}

type Servicio struct {
	nombre string
	precio float64
	min    uint
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(prods []Producto, c chan float64) {
	total := 0.0
	for _, prod := range prods {
		total += (prod.precio * float64(prod.cant))
	}
	c <- total
}

func sumarServicios(servs []Servicio, c chan float64) {
	total := 0.0
	var cantTrabajo int
	for _, serv := range servs {
		if serv.min <= 30 {
			cantTrabajo = 1
		} else {
			cantTrabajo = int(serv.min) / 30
		}
		total += (float64(cantTrabajo) * serv.precio)
	}
	c <- total
}

func sumarMantenimiento(mants []Mantenimiento, c chan float64) {
	total := 0.0
	for _, mant := range mants {
		total += mant.precio
	}
	c <- total
}

func main() {
	p := Producto{"PC", 60000, 2}
	s := Servicio{"Internet", 1000, 30}
	m := Mantenimiento{"Mantenimiento internet", 500}

	ps := []Producto{p}
	ss := []Servicio{s}
	ms := []Mantenimiento{m}

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go sumarProductos(ps, c1)
	go sumarServicios(ss, c2)
	go sumarMantenimiento(ms, c3)

	total := <-c1 + <-c2 + <-c3

	fmt.Println("El total de todo es: ", total)
}
