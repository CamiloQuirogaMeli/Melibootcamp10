package main

import "fmt"

const (
	PEQUENIO = "pequenio"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
)

type Ecommerce interface {
	precio() float64
	envio() float64
}

// TIENDA 1
type Store1 struct {
	nombre string
}

func (s1 Store1) precio() float64 {
	return //retorna precio
}

func (s1 Store1) envio() string {
	return //retorna envio
}

// TIENDA 2
type Store2 struct {
	nombre string
}

func (s1 Store2) precio() float64 {
	return //retorna precio
}

func (s1 Store2) envio() string {
	return //retorna envio
}

func nuevaTienda(prodType string) Ecommerce {
	switch prodType {
	case PEQUENIO:
		return Store1{}
	}
}

func main() {
	fmt.Println("vim-go")
	p := nuevaTienda(PEQUENIO)
	m := nuevaTienda(MEDIANO)
	g := nuevaTienda(GRANDE)

	fmt.Println(p.precio())
	fmt.Println(p.envio())
	fmt.Println(m.precio())
	fmt.Println(m.envio())
	fmt.Println(g.precio())
	fmt.Println(g.envio())
}
