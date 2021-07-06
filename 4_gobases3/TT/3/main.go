package main

import "fmt"

func main() {

	productos := []Producto{
		{
			nombre:   "remera",
			precio:   500.00,
			cantidad: 5,
		},
		{
			nombre:   "buzo",
			precio:   1000.00,
			cantidad: 2,
		}}

	servicios := []Servicio{
		{
			nombre:            "servicio1",
			precio:            100.00,
			minutosTrabajados: 25,
		},
		{
			nombre:            "servicio2",
			precio:            200.00,
			minutosTrabajados: 60,
		}}

	mantenimientos := []Mantenimiento{
		{
			nombre: "mantenimiento1",
			precio: 1000.00,
		},
		{
			nombre: "mantenimiento2",
			precio: 1500.00,
		}}

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)
	go SumarProductos(productos, c1)
	go SumarServicios(servicios, c2)
	go SumarMantenimiento(mantenimientos, c3)

	totalProductos := <-c1
	totalServicios := <-c2
	totalMantenimientos := <-c3
	total := totalProductos + totalServicios + totalMantenimientos

	fmt.Println("Suma total: ", total)
}

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

func SumarProductos(productos []Producto, c chan float64) {
	var monto float64
	for _, p := range productos {
		monto += (p.precio * float64(p.cantidad))
	}
	c <- monto
}
func SumarServicios(servicios []Servicio, c chan float64) {
	var monto float64
	for _, s := range servicios {
		if s.minutosTrabajados < 30 {
			monto += s.precio / 2
		} else {
			monto += s.precio
		}
	}
	c <- monto
}
func SumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	var monto float64
	for _, m := range mantenimientos {
		monto += m.precio
	}
	c <- monto
}
