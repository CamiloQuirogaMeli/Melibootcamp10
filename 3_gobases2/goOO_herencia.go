package main

import "fmt"

//GO ORIENTADO A OBJETOS - HERENCIA

type Vehiculo struct {
	km     float64
	tiempo float64
}
type Auto struct {
	v Vehiculo
}
type Moto struct {
	v Vehiculo
}

func (v Vehiculo) detalle() {
	fmt.Printf("km: \t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}

func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

/*
func mainHerencia() {
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(120)
	moto.Detalle()
}
*/
