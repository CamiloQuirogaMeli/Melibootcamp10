package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}
type Servicio struct {
	Nombre            string
	Precio            float32
	MinutosTrabajados int
}
type Mantenimiento struct {
	Nombre string
	Precio float32
}

func main() {
	var sumaTotal float32 = 0.0
	listaProductos := []Producto{
		{
			Nombre:   "Producto 1",
			Precio:   15.8,
			Cantidad: 20,
		},
		{
			Nombre:   "Producto 2",
			Precio:   100.3,
			Cantidad: 10,
		},
		{
			Nombre:   "Producto 3",
			Precio:   50.8,
			Cantidad: 100,
		},
	}
	listaServicios := []Servicio{
		{
			Nombre:            "Servicio 1",
			Precio:            100.50,
			MinutosTrabajados: 90,
		},
		{
			Nombre:            "Servicio 2",
			Precio:            530.25,
			MinutosTrabajados: 45,
		},
		{
			Nombre:            "Servicio 3",
			Precio:            95.48,
			MinutosTrabajados: 160,
		},
	}
	listaMantenimientos := []Mantenimiento{
		{
			Nombre: "Mantenimiento 1",
			Precio: 500.32,
		},
		{
			Nombre: "Mantenimiento 2",
			Precio: 1000.92,
		},
		{
			Nombre: "Mantenimiento 3",
			Precio: 532.82,
		},
	}
	canalProductos, canalServicios, canalMantenimientos := make(chan float32), make(chan float32), make(chan float32)

	go SumarProductos(listaProductos, canalProductos)
	go SumarServicios(listaServicios, canalServicios)
	go SumarMantenimientos(listaMantenimientos, canalMantenimientos)
	sumaTotal = <-canalProductos + <-canalServicios + <-canalMantenimientos
	fmt.Println("La sumatoria total es: ", sumaTotal)
}

func SumarProductos(productos []Producto, c chan float32) {
	var sumatoria float32 = 0.0
	for _, prod := range productos {
		if prod.Cantidad >= 0 && prod.Precio >= 0.0 {
			sumatoria += float32(prod.Cantidad) * prod.Precio
		} else {
			continue
		}
	}
	c <- sumatoria
}
func SumarServicios(servicios []Servicio, c chan float32) {
	var sumatoria float32 = 0.0
	for _, serv := range servicios {
		if serv.MinutosTrabajados >= 0 && serv.Precio >= 0.0 {
			if serv.MinutosTrabajados <= 30 {
				sumatoria += serv.Precio * 30
			} else {
				sumatoria += serv.Precio * float32(serv.MinutosTrabajados/30)
			}
		} else {
			continue
		}
	}
	c <- sumatoria
}
func SumarMantenimientos(mantenimientos []Mantenimiento, c chan float32) {
	var sumatoria float32 = 0.0
	for _, mant := range mantenimientos {
		if mant.Precio >= 0 {
			sumatoria += mant.Precio
		} else {
			continue
		}
	}
	c <- sumatoria
}
