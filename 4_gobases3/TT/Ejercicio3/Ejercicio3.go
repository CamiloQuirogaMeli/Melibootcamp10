package main

import "fmt"

func main() {
	/*
		Ejercicio 3 - Calcular Precio
		Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
		Para ello requieren realizar un programa que se encargue de calcular el precio total de
		Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la
		velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go
		routines.

		Se requieren 3 estructuras:
		- Productos: nombre, precio, cantidad.
		- Servicios: nombre, precio, minutos trabajados.
		- Mantenimiento: nombre, precio.

		Se requieren 3 funciones:
		- Sumar Productos: recibe un array de producto y devuelve el precio total (precio *
		cantidad).
		- Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media
		hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado
		media hora).
		- Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

		Los 3 se deben ejecutar en paralelo y al final se debe mostrar por pantalla el monto final
		(sumando el total de los 3).
	*/
	p1 := Product{"Producto1", 500, 4}
	p2 := Product{"Producto2", 4500, 2}
	p3 := Product{"Producto3", 1600, 3}
	s1 := Service{"Service1", 200, 60}
	s2 := Service{"Service2", 230, 25}
	s3 := Service{"Service3", 550, 120}
	m1 := Maintenance{"Maintenance1", 300}
	m2 := Maintenance{"Maintenance2", 230}
	m3 := Maintenance{"Maintenance3", 450}

	p := []Product{p1, p2, p3}
	s := []Service{s1, s2, s3}
	m := []Maintenance{m1, m2, m3}

	chanel1 := make(chan float64)
	chanel2 := make(chan float64)
	chanel3 := make(chan float64)

	go addProduct(p, chanel1)
	go addService(s, chanel2)
	go addMaintenance(m, chanel3)

	total := <-chanel1 + <-chanel2 + <-chanel3

	fmt.Println("Total:", total)

}

type Product struct {
	name  string
	price float64
	cant  int
}
type Service struct {
	name    string
	price   float64
	minWork float64
}
type Maintenance struct {
	name  string
	price float64
}

func addProduct(products []Product, c chan float64) {
	var total float64 = 0

	for _, p := range products {
		total += p.price * float64(p.cant)
	}
	c <- total
}

func addService(services []Service, c chan float64) {
	var total float64
	var cant int
	for _, s := range services {
		if s.minWork <= 30 {
			cant = 1
		} else {
			cant = int(s.minWork) / 30
		}
		total += float64(cant) * (s.price)
	}
	c <- total
}

func addMaintenance(maintenances []Maintenance, c chan float64) {
	var total float64 = 0
	for _, m := range maintenances {
		total += m.price
	}
	c <- total
}
