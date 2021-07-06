package main

import (
	"fmt"
	"math"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

type geometry interface { //implementa las funciones propias de cada objeto
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { //recibe un objeto de algun tipo e imprime los resultados de las funciones de cada objeto
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func mainInterfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) //le paso el objeto rectangulo
	measure(c) //le paso el objeto circulo

	//OTRO TIPO PARA CREAR OBJETO
	rectangulo := newGeometry(rectType, 2, 3)
	measure(rectangulo)
	circulo := newGeometry(circleType, 6)
	measure(circulo)

	//CASTEO DE VARIABLES
	var suma int = 100
	var importe int = 19
	var promedio float32

	promedio = float32(suma) / float32(importe)
	fmt.Println("El promedio es: ", promedio)

	//ASERCION DE TIPO DE DATOS
	var i interface{} = "hola"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

}
