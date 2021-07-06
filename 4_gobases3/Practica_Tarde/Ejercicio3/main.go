package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre        string
	precio        float64
	min_trabajado int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProducto(pr []Producto, c chan float64) {
	var precio float64
	for _, prod := range pr {
		precio += prod.precio * float64(prod.cantidad)
	}
	c <- precio
}

func sumarServ(ser []Servicio, c chan float64) {
	var precio float64
	for _, serv := range ser {
		if serv.min_trabajado < 30 {
			precio += serv.precio * 30
		} else {
			precio += serv.precio * float64(serv.min_trabajado)
		}
	}
	c <- precio
}

func sumarMantenimiento(man []Mantenimiento, c chan float64) {
	var precio float64
	for _, mant := range man {
		precio += mant.precio
	}
	c <- precio
}

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	p1 := Producto{nombre: "Producto1", precio: 200, cantidad: 3}       //600
	s1 := Servicio{nombre: "Producto1", precio: 200, min_trabajado: 40} //8000
	m1 := Mantenimiento{nombre: "Producto1", precio: 200}               //200

	ps := []Producto{p1}
	ss := []Servicio{s1}
	ms := []Mantenimiento{m1}

	go sumarProducto(ps, c1)
	go sumarServ(ss, c2)
	go sumarMantenimiento(ms, c3)

	montoTotal := <-c1 + <-c2 + <-c3

	fmt.Println("El monto final es: ", montoTotal)

}
