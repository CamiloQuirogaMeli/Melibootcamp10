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
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(productos []Producto, c chan float64) (precioTotal float64) {
	for _, value := range productos {
		precioTotal += value.Precio
	}
	c <- precioTotal
	return
}

func SumarServicios(servicios []Servicio, c chan float64) (precioTotal float64) {
	for _, value := range servicios {
		valorMinuto := value.Precio / float64(30)
		if value.Minutos < 30 {
			precioTotal += valorMinuto * 30
		} else {
			precioTotal += valorMinuto * float64(value.Minutos)
		}
	}
	c <- precioTotal
	return
}

func SumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) (precioTotal float64) {
	for _, value := range mantenimientos {
		precioTotal += value.Precio
	}
	c <- precioTotal
	return
}

func main() {
	p1 := Producto{"Coca-Cola", 158.5, 200}
	p2 := Producto{"Pepsi", 147.5, 100}
	p3 := Producto{"Seven-up", 151.5, 80}
	sliceP := []Producto{p1, p2, p3}

	s1 := Servicio{"Limpieza", 280, 240}
	s2 := Servicio{"Limpieza", 275, 300}
	s3 := Servicio{"Limpieza", 300, 20}
	sliceS := []Servicio{s1, s2, s3}

	m1 := Mantenimiento{"Mante", 500}
	m2 := Mantenimiento{"Mante", 600}
	sliceM := []Mantenimiento{m1, m2}

	c := make(chan float64)
	c1 := make(chan float64)
	c2 := make(chan float64)

	go SumarProductos(sliceP, c)
	go SumarServicios(sliceS, c1)
	go SumarMantenimiento(sliceM, c2)

	fmt.Println("El precio total es: ", <-c+<-c1+<-c2)

}
