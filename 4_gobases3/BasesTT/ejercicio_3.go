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
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func SumarProducto(productos []Producto, c chan float64) {
	var total float64
	fmt.Println("Suma de productos")
	for _, p := range productos {
		total += float64(p.cantidad) * p.precio
	}
	c <- total
}

func SumarServicio(servicios []Servicio, c chan float64) {
	var total float64
	fmt.Println("Suma de servicios")
	for _, s := range servicios {
		total += s.precio * math.Ceil(float64(s.minutosTrabajados)/30.0)
	}
	c <- total
}

func SumarManenimiento(mantenimientos []Mantenimiento, c chan float64) {
	var total float64
	fmt.Println("Suma de mantenimientos")
	for _, m := range mantenimientos {
		total += m.precio
	}
	c <- total
}

func main() {
	canalSuma := make(chan float64)
	productos := []Producto{Producto{"p1", 10.0, 2}, Producto{"p2", 20.0, 1}}
	fmt.Println("Productos:", productos)
	servicios := []Servicio{Servicio{"s1", 10.0, 29}, Servicio{"s2", 20.0, 61.0}}
	fmt.Println("Servicios:", servicios)
	mantenimientos := []Mantenimiento{Mantenimiento{"m1", 23.0}, Mantenimiento{"m2", 27.0}}
	fmt.Println("Mantenimientos:", mantenimientos)
	go SumarProducto(productos, canalSuma)
	go SumarServicio(servicios, canalSuma)
	go SumarManenimiento(mantenimientos, canalSuma)
	a, b, c := <-canalSuma, <-canalSuma, <-canalSuma
	fmt.Println("Total:", a+b+c)
}
