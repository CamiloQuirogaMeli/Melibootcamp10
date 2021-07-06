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

package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  float64
	Minutos float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos []Producto, total chan float64) {
	var totalProductos float64
	for _, p := range productos {
		total := p.Precio * float64(p.Cantidad)
		println("El precio total del producto", p.Nombre, "es:", total)
		totalProductos += total
	}
	total <- totalProductos
}

func sumarServicios(servicios []Servicio, total chan float64) {
	var totalServicios float64
	for _, s := range servicios {
		if s.Minutos < 30 {
			s.Minutos = 30
		}
		total := s.Minutos * s.Precio
		totalServicios += total
	}
	total <- totalServicios
}

func sumarMatenimiento(mantenimientos []Mantenimiento, total chan float64) {
	var totalMantenimientos float64
	for _, m := range mantenimientos {
		totalMantenimientos += m.Precio
	}
	total <- totalMantenimientos
}

func main() {
	producto := Producto{Nombre: "Product", Precio: 10.99, Cantidad: 3}
	var ProductoSlice []Producto
	servicio := Servicio{Nombre: "Servicio", Precio: 30.15, Minutos: 10}
	var ServicioSlice []Servicio
	mantenimiento := Mantenimiento{Nombre: "Mantenimiento", Precio: 11.30}
	var MantenimientoSlice []Mantenimiento
	for i := 0; i < 5; i++ {
		ProductoSlice = append(ProductoSlice, producto)
		ServicioSlice = append(ServicioSlice, servicio)
		MantenimientoSlice = append(MantenimientoSlice, mantenimiento)
	}
	totalProductos := make(chan float64)
	totalServicios := make(chan float64)
	totalMantenimiento := make(chan float64)

	go sumarProductos(ProductoSlice, totalProductos)
	go sumarServicios(ServicioSlice, totalServicios)
	go sumarMatenimiento(MantenimientoSlice, totalMantenimiento)

	TotalFinal := <-totalProductos + <-totalServicios + <-totalMantenimiento

	fmt.Println("Total:", TotalFinal)
}
