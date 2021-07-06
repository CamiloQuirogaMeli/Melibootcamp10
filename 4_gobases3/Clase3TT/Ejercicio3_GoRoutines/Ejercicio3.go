// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de
// Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la

// 3
// velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go
// routines.

// Se requieren 3 estructuras:
// - Productos: nombre, precio, cantidad.
// - Servicios: nombre, precio, minutos trabajados.
// - Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// - Sumar Productos: recibe un array de producto y devuelve el precio total (precio *
// cantidad).
// - Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media
// hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado
// media hora).
// - Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar en paralelo y al final se debe mostrar por pantalla el monto final
// (sumando el total de los 3).

package main

import (
	"fmt"
)

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type Service struct {
	name    string
	price   float64
	minutes int
}

type Maintenance struct {
	Name  string
	Price float64
}

func totalProduct(prod []Product, c chan float64) {
	var total float64

	for _, p := range prod {
		total += p.Price * float64(p.Amount)
	}
	c <- total
}

func totalMaintenance(main []Maintenance, c chan float64) {
	var total float64

	for _, m := range main {
		total += m.Price
	}
	c <- total
}

func totalService(serv []Service, c chan float64) {
	var total float64
	for _, s := range serv {
		if s.minutes < 30 {
			total += s.price
		} else {
			total += s.price * (float64(s.minutes) / 30)
		}
	}
	c <- total
}

func main() {

	Products := []Product{{"Zapatos", 199.0, 5}, {"Camisa", 99.9, 10}, {"Pantalon", 150.9, 2}}
	Services := []Service{{"Build a House", 100.0, 3000}, {"Web Application", 300.0, 28}}
	Maintenances := []Maintenance{{"Computing", 55.99}, {"Cleaning", 199.0}}

	c := make(chan float64)
	b := make(chan float64)
	a := make(chan float64)
	go totalProduct(Products, c)
	go totalMaintenance(Maintenances, b)
	go totalService(Services, a)
	sumaTotal := <-c + <-b + <-a

	fmt.Println("El valor total es de: %f", sumaTotal)
}
