package main

import (
	"errors"
	"fmt"
)

const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type Ecommerce interface {
	precio() float64
	envio() string
}

type Producto struct {
	valor float64
	direc string
}

type Pequenio struct {
	p Producto
}

type Mediano struct {
	p Producto
}

type Grande struct {
	p Producto
}

func (p Pequenio) precio() float64 {
	return p.p.valor
}

func (m Mediano) precio() float64 {
	return m.p.valor + (m.p.valor * 0.03)
}

func (g Grande) precio() float64 {
	return g.p.valor + (g.p.valor * 0.06) + 2500
}

func (c Pequenio) envio() string {
	return c.p.direc
}

func (m Mediano) envio() string {
	return m.p.direc
}

func (g Grande) envio() string {
	return g.p.direc
}

func detalles(e Ecommerce) {
	fmt.Printf("Precio: %f\nDireccion: %s\n", e.precio(), e.envio())
}

func nuevaTienda(tipo string, monto float64, direccion string) (Ecommerce, error) {
	switch tipo {
	case CHICO:
		return Pequenio{Producto{monto, direccion}}, nil
	case MEDIANO:
		return Mediano{Producto{monto, direccion}}, nil
	case GRANDE:
		return Grande{Producto{monto, direccion}}, nil
	}
	return nil, errors.New("Error al ingresar el tipo de producto")

}

func main() {
	p1, er := nuevaTienda(CHICO, 100, "Diego Diaz")
	if er != nil {
		fmt.Print(er)
	} else {
		detalles(p1)
	}
	fmt.Println("-----------")

	p2, er := nuevaTienda(MEDIANO, 200, "Leopoldo Buteler")
	if er != nil {
		fmt.Print(er)
	} else {
		detalles(p2)
	}
	fmt.Println("-----------")

	p3, er := nuevaTienda(GRANDE, 300, "Alonso de Ubeda")
	if er != nil {
		fmt.Print(er)
	} else {
		detalles(p3)
	}

}
