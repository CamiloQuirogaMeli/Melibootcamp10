package main

import (
	"errors"
	"fmt"
)

type Chico struct {
	precio float64
}

type Mediano struct {
	precio float64
}

type Grande struct {
	precio float64
}

const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

func (c *Chico) setPrecio(costo float64) {
	c.precio = costo
}

func (c Chico) getPrecio() float64 {
	return c.precio
}

func (m *Mediano) setPrecio(costo float64) {
	m.precio = costo
}

func (m Mediano) getPrecio() float64 {
	return m.precio + (m.precio * 0.03)
}

func (g *Grande) setPrecio(costo float64) {
	g.precio = costo
}

func (g Grande) getPrecio() float64 {
	return g.precio + (g.precio * 0.06) + 2500
}

type Producto interface {
	getPrecio() float64
}

func factory(tipoProd string, precio float64) (Producto, error) {
	switch tipoProd {
	case CHICO:
		{
			var c Chico
			c.setPrecio(precio)
			return c, nil
		}
	case MEDIANO:
		{
			var m Mediano
			m.setPrecio(precio)
			return m, nil
		}
	case GRANDE:
		{
			var g Grande
			g.setPrecio(precio)
			return g, nil
		}
	default:
		{
			return nil, errors.New("lo sentimos mucho, aun no disponemos de ese producto")
		}
	}
}

func main() {
	var (
		tipoProd string
		precio   float64
	)
	fmt.Println("Ingrese cual producto desea consultar el costo (chico, mediano, grande)")
	fmt.Scanln(&tipoProd)
	fmt.Println("Ingrese el precio del producto")
	fmt.Scanln(&precio)
	if precio <= 0 {
		fmt.Println("El precio no puede ser menor o igual a 0")
	} else {
		producto, err := factory(tipoProd, precio)
		if err != nil {
			fmt.Println(err)
		} else {
			costo := producto.getPrecio()
			fmt.Println("El costo total sera de: ", costo)
		}
	}
}
