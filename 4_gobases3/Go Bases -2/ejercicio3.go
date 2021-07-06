package main

import "fmt"

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

type Producto struct {
	nombre   string
	precio   int
	cantidad int
}

type Servicio struct {
	nombre  string
	precio  int
	minutos int
}

type Mantenimiento struct {
	nombre string
	precio int
}

func sumarProductos(canal1 chan int, sliceP []Producto) float64 {
	var sumaP = 0
	for _, p := range sliceP {
		sumaP += p.precio * p.cantidad
	}
	//fmt.Println(sumaP)
	canal1 <- sumaP
	return float64(sumaP)
}

func sumarServicios(canal2 chan int, sliceS []Servicio) float64 {
	var sumaS = 0
	for _, s := range sliceS {
		sumaS += s.precio * s.minutos
	}
	//fmt.Println(sumaS)
	canal2 <- sumaS
	return float64(sumaS)

}

func sumarMantenimientos(canal3 chan int, sliceM []Mantenimiento) float64 {
	var sumaM = 0
	for _, m := range sliceM {
		sumaM += m.precio
	}
	//fmt.Println(sumaM)
	canal3 <- sumaM
	return float64(sumaM)
}

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	canal3 := make(chan int)

	productos := []Producto{
		{
			nombre:   "arroz",
			precio:   1000,
			cantidad: 3,
		},
		{
			nombre:   "azucar",
			precio:   1200,
			cantidad: 2,
		},
		{
			nombre:   "sal",
			precio:   800,
			cantidad: 4,
		},
	}

	servicios := []Servicio{
		{
			nombre:  "cocinar",
			precio:  1000,
			minutos: 2,
		},
		{
			nombre:  "barrer",
			precio:  1200,
			minutos: 4,
		},
		{
			nombre:  "trapear",
			precio:  1000,
			minutos: 2,
		},
	}

	mantenimientos := []Mantenimiento{
		{
			nombre: "preventivo",
			precio: 1000,
		},
		{
			nombre: "correctivo",
			precio: 1200,
		},
		{
			nombre: "normal",
			precio: 1000,
		},
	}

	go sumarProductos(canal1, productos)
	go sumarServicios(canal2, servicios)
	go sumarMantenimientos(canal3, mantenimientos)

	resultadoP := <-canal1
	fmt.Println("Total de productos", resultadoP)
	resultadoS := <-canal2
	fmt.Println("Total de servicios", resultadoS)
	resultadoM := <-canal3
	fmt.Println("Total de mantenimientos", resultadoM)

	fmt.Println("La suma total es: ", (resultadoM + resultadoP + resultadoS))
}
