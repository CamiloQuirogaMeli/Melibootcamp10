package main

import "fmt"

/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de
Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la
velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

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

type ProductsCP struct {
	NameP   string
	CostP   float64
	AmountP int
}

type ServicesCP struct {
	NameServ string
	CostServ float64
	Minutes  float64
}

type ManteinmentCP struct {
	NameMant string
	CostMant float64
}

func addProducts(p []ProductsCP, c chan float64) (totalCost float64) {
	var sumatoria float64
	for _, value := range p {
		sumatoria += value.CostP
		totalCost = sumatoria * float64(value.AmountP)
	}
	fmt.Println("Fin suma de Productos", totalCost)
	c <- totalCost
	return
}

func addServices(s []ServicesCP, c chan float64) float64 {
	var duracion float64
	var totalCostS float64
	for _, value := range s {
		duracion = value.Minutes / 30
		if duracion < 1 {
			totalCostS += value.CostServ * 1
		} else {
			totalCostS += duracion * value.CostServ
		}
	}
	fmt.Println("Fin suma de Servicios", totalCostS)
	c <- totalCostS
	return totalCostS
}

func addManteinance(m []ManteinmentCP, c chan float64) (totalCostM float64) {
	for _, maint := range m {
		totalCostM += maint.CostMant
	}
	fmt.Println("Fin suma de Mantenimiento", totalCostM)
	c <- totalCostM
	return
}
