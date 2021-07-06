package main

import (
	"fmt"
	"math"
)

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre        string
	precio        float64
	minTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProds(c chan float64, products ...Producto) {
	var suma float64
	for _, p := range products {
		suma += p.precio * float64(p.cantidad)
	}
	fmt.Println("Suma productos ", suma)
	c <- suma
}

func sumarServicio(c chan float64, services ...Servicio) {
	var suma float64
	for _, s := range services {
		suma += s.precio * math.Ceil(float64(s.minTrabajados)/30.0)
	}
	fmt.Println("Suma servicios ", suma)
	c <- suma
}

func sumarMantenimiento(c chan float64, maintenece ...Mantenimiento) {
	var suma float64
	for _, m := range maintenece {
		suma += m.precio
	}
	fmt.Println("Suma mantenimientos ", suma)
	c <- suma
}

func main() {
	a := make(chan float64)
	b := make(chan float64)
	c := make(chan float64)
	prod1 := Producto{
		nombre:   "banana",
		precio:   30,
		cantidad: 5,
	}
	prod2 := Producto{
		nombre:   "manzana",
		precio:   15.60,
		cantidad: 3,
	}
	serv1 := Servicio{
		nombre:        "reparar PC",
		precio:        25,
		minTrabajados: 220,
	}
	serv2 := Servicio{
		nombre:        "pintar habitacion",
		precio:        80,
		minTrabajados: 400,
	}
	man1 := Mantenimiento{
		nombre: "mantenimiento 1",
		precio: 700,
	}
	man2 := Mantenimiento{
		nombre: "mantenimiento 2",
		precio: 500,
	}
	go sumarProds(a, prod1, prod2)
	go sumarServicio(b, serv1, serv2)
	go sumarMantenimiento(c, man1, man2)
	sumatoria := <-a
	sumatoria += <-b
	sumatoria += <-c
	fmt.Print("Suma total ", sumatoria)

}
