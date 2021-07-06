/* Ejercicio 3 - Calcular Precio
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
(sumando el total de los 3). */
package main

import (
	"fmt"
	"math"
)

type Productos struct {
	nombre string
	precio float64
	cantidad int
}

type Servicios struct {
	nombre string
	precio float64
	minutos int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(arr []Productos, c chan float64) float64 {
	total := 0.0
	for _, val := range arr {
		total += val.precio * float64(val.cantidad)
	}
	c <- total
	return total
}

func sumarServicios(arr []Servicios, c chan float64) float64 {
	total := 0.0
	for i := 0; i < len(arr); i++ {
		minutos := math.Ceil(float64(arr[i].minutos / 30))
		total += arr[i].precio * minutos
	}
	c <- total
	return total
}

func sumarMantenimiento(arr []Mantenimiento, c chan float64) float64 {
	total := 0.0
	for _, val := range arr {
		total += val.precio
	}
	c <- total
	return total
}

var testProductos []Productos
var testServicios []Servicios
var testMantenimiento []Mantenimiento

func main() {
	fmt.Println("Empezo el programa")

	c := make(chan float64)
	testProductos = append(testProductos, Productos{"test", 100.0, 3}, Productos{"test2", 250.0, 5})
	testServicios = append(testServicios, Servicios{"test", 100.0, 45}, Servicios{"test2", 250.0, 60})
	testMantenimiento = append(testMantenimiento, Mantenimiento{"test", 500.0}, Mantenimiento{"test2", 450.0})

	go sumarProductos(testProductos, c)
	go sumarServicios(testServicios, c)
	go sumarMantenimiento(testMantenimiento, c)
	total := 0.0
	total += <-c
	total += <-c
	total += <-c
	fmt.Println("Total:", total)
	fmt.Println("Terminó el programa")
}