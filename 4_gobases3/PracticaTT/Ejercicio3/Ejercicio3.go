// Ejercicio 3 - Calcular Precio
// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
// Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar en paralelo y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

package main

import (
	f "fmt"
)

type Product struct {
	name  string
	price float64
	cant  int
}

type Service struct {
	name        string
	price       float64
	minutesWork int
}

type Maintenance struct {
	name  string
	price float64
}

func sumProducts(products []Product, c chan float64) {
	var totalPrice float64 = 0

	for _, prod := range products {
		totalPrice += prod.price * float64(prod.cant)
	}
	c <- totalPrice
}

func sumServices(services []Service, c chan float64) {
	var totalPrice float64 = 0

	for _, serv := range services {
		if serv.minutesWork%30 != 0 {
			totalPrice += float64(serv.minutesWork * (serv.minutesWork/30 + 1))
		} else {
			totalPrice += float64(serv.minutesWork * (serv.minutesWork / 30))
		}
	}
	c <- totalPrice
}

func sumMaintenance(maintenances []Maintenance, c chan float64) {
	var totalPrice float64 = 0

	for _, mainten := range maintenances {
		totalPrice += mainten.price
	}
	c <- totalPrice
}

func main() {
	c := make(chan float64)
	var sum float64 = 0

	var prod1 Product = Product{"fideos", 120.60, 3}
	var prod2 Product = Product{"dulce de leche", 140.75, 20}
	var prodSlice []Product = []Product{prod1, prod2}
	var mainten1 Maintenance = Maintenance{"computadora", 700}
	var mainten2 Maintenance = Maintenance{"servidores", 45043}
	var maintenSlice []Maintenance = []Maintenance{mainten1, mainten2}
	var service1 Service = Service{"limpieza", 4563, 4023}
	var service2 Service = Service{"charter", 3409, 35}
	var serviceSlice []Service = []Service{service1, service2}

	go sumMaintenance(maintenSlice, c)
	go sumServices(serviceSlice, c)
	go sumProducts(prodSlice, c)

	sum += <-c
	sum += <-c
	sum += <-c

	f.Printf("La suma de servicios, productos y mantenimiento es: %.2f\n", sum)
}
