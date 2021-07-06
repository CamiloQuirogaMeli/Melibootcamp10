package main

import "fmt"

/*
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

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type servicio struct {
	nombre        string
	precio        float64
	minTrabajados int
}

type mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(productos []producto, c chan float64) {
	sumProductos := 0.0

	for _, p := range productos {
		sumProductos += float64(p.cantidad) * (p.precio)
	}
	fmt.Println("Productos: ", sumProductos)
	c <- sumProductos
}

func sumarServicios(servicios []servicio, c chan float64) {
	sumServicios := 0.0

	for _, s := range servicios {
		if s.minTrabajados < 30 {
			sumServicios += s.precio
		} else {
			sumServicios += s.precio * (float64(s.minTrabajados) / float64(30))
		}
	}
	fmt.Println("Servicios: ", sumServicios)
	c <- sumServicios
}

func sumarMantenimientos(mantenimientos []mantenimiento, c chan float64) {
	sumMantenimientos := 0.0

	for _, m := range mantenimientos {
		sumMantenimientos += m.precio
	}
	fmt.Println("Mantenimientos: ", sumMantenimientos)
	c <- sumMantenimientos
}

func main() {

	/*Estructuras*/
	p1 := producto{"Mouse", 1000.0, 2}
	p2 := producto{"Teclado", 1500.0, 1}
	s1 := servicio{"Internet", 50.0, 20}
	s2 := servicio{"Luz", 15.0, 105}
	m1 := mantenimiento{"Arreglo pc", 1000.0}
	m2 := mantenimiento{"Instalaciones", 500.0}

	/*Declaro los canales*/
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	var p []producto
	var s []servicio
	var m []mantenimiento

	p = append(p, p1, p2)
	s = append(s, s1, s2)
	m = append(m, m1, m2)

	go sumarProductos(p, c1)
	go sumarServicios(s, c2)
	go sumarMantenimientos(m, c3)

	sumaTotal := <-c1 + <-c2 + <-c3
	fmt.Println("Suma total: ", sumaTotal)

	fmt.Println()
}
