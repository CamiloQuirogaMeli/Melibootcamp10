package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := producto{nombre: "Pinza", precio: 980.00, cantidad: 2}
	p2 := producto{nombre: "Destornillador pala", precio: 420.00, cantidad: 3}
	s1 := servicio{nombre: "Instalacion Aire", precio: 2500.00, minutosTrabajados: 120}
	s2 := servicio{nombre: "Instalacion cable", precio: 1500.00, minutosTrabajados: 60}
	m1 := mantenimiento{nombre: "Revision calefactores", precio: 2000.00}
	m2 := mantenimiento{nombre: "Limpieza filtros aire", precio: 1500.00}
	productos := []producto{p1, p2}
	servicios := []servicio{s1, s2}
	mantenimientos := []mantenimiento{m1, m2}

	channel1 := make(chan float64)
	channel2 := make(chan float64)
	channel3 := make(chan float64)

	go sumarProductos(productos, channel1)
	go sumarServicios(servicios, channel2)
	go sumarMantenimiento(mantenimientos, channel3)

	montoTotal := <-channel1 + <-channel2 + <-channel3

	fmt.Println("El monto total es de:", montoTotal)

	// SOLUCION SIN CHANNELS
	// slistaCompletaVentasServiciosMantenimientos := []float64{sumarProductos(p1, p2), sumarServicios(s1, s2), sumarMantenimiento(m1, m2)}
	// var montoTotal float64
	// for _, val := range slistaCompletaVentasServiciosMantenimientos {
	// 	montoTotal += val
	// }
	//fmt.Println(montoTotal)

	// fmt.Println("El total de los productos es:", sumarProductos(p1, p2))
	// fmt.Println("El total de los servicios es:", sumarServicios(s1, s2))
	// fmt.Println("El total de los mantenimientos es:", sumarMantenimiento(m1, m2))

}

type producto struct {
	nombre   string
	precio   float64
	cantidad int16
}

type servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int32
}

type mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(productos []producto, c chan float64) {
	precioFinalProductos := 0.00
	for i := 0; i < len(productos); i++ {
		precioFinalProductos += productos[i].precio * float64(productos[i].cantidad)
	}
	c <- precioFinalProductos
}

func sumarServicios(servicios []servicio, c chan float64) {
	precioFinalServicios := 0.00
	for i := 0; i < len(servicios); i++ {
		precioFinalServicios += servicios[i].precio * math.RoundToEven(float64(servicios[i].minutosTrabajados/30.0))
	}
	c <- precioFinalServicios
}

func sumarMantenimiento(mantenimientos []mantenimiento, c chan float64) {
	precioFinalMantenimientos := 0.00
	for i := 0; i < len(mantenimientos); i++ {
		precioFinalMantenimientos += mantenimientos[i].precio
	}
	c <- precioFinalMantenimientos
}
