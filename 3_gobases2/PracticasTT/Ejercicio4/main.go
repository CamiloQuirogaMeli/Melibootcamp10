package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

const (
	CHICO    = "chico"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
	ESPECIAL = "especial"
	FRAGIL   = "fragil"
	LIMITE   = 10000000.0
)

type Chico struct {
	size float64
}

type Mediano struct {
	size float64
}

type Grande struct {
	size float64
}

type Especial struct {
	size float64
}

type Fragil struct {
	size float64
}

func (c *Chico) setSize(espacio float64) {
	c.size = espacio
}

func (c Chico) getSize() float64 {
	return c.size
}

func (m *Mediano) setSize(espacio float64) {
	m.size = espacio
}

func (m Mediano) getSize() float64 {
	return m.size + (m.size * 0.05)
}

func (g *Grande) setSize(espacio float64) {
	g.size = espacio
}

func (g Grande) getSize() float64 {
	return g.size + (g.size * 0.20)
}

func (f *Fragil) setSize(espacio float64) {
	f.size = espacio
}

func (f Fragil) getSize() float64 {
	return f.size + (f.size * 0.75)
}

func (e *Especial) setSize(espacio float64) {
	e.size = espacio
}

func (e Especial) getSize() float64 {
	return e.size
}

type Flete struct {
	productos []Producto
}

type Producto interface {
	getSize() float64
}

func (F *Flete) AgregarProducto(p Producto) {
	F.productos = append(F.productos, p)
}

func CrearProducto(peso float64, tipoProd string) (Producto, error) {
	switch tipoProd {
	case CHICO:
		{
			var c Chico
			c.setSize(peso)
			return c, nil
		}
	case MEDIANO:
		{
			var m Mediano
			m.setSize(peso)
			return m, nil
		}
	case GRANDE:
		{
			var g Grande
			g.setSize(peso)
			return g, nil
		}
	case FRAGIL:
		{
			var f Fragil
			f.setSize(peso)
			return f, nil
		}
	case ESPECIAL:
		{
			var e Especial
			e.setSize(peso)
			return e, nil
		}
	default:
		{
			return nil, errors.New("no existe este tipo de producto")
		}
	}
}

func (F Flete) CalcularEnvios() int {
	var espacio, espacioEsp float64
	for _, product := range F.productos {
		if reflect.TypeOf(product).Name() == "Especial" {
			espacioEsp += product.getSize()
		} else {
			espacio += product.getSize()
		}
	}
	enviosNorm := (espacio / LIMITE)
	enviosEsp := (espacioEsp / LIMITE)
	envios := math.Ceil(enviosNorm) + math.Ceil(enviosEsp)
	return int(envios)
}

func main() {
	bucle := true
	var (
		peso     float64
		tipoProd string
		flete    Flete
	)
	for bucle {
		fmt.Println("Ingrese un tipo de producto (chico, mediano, grande, fragil, especial) para agregar al flete")
		fmt.Scanln(&tipoProd)
		fmt.Println("Ingrese el espacio que ocupa este producto en cm3")
		fmt.Scanln(&peso)
		if peso <= 0 {
			fmt.Println("Dato de espacio incorrecto")
		} else {
			prod, err := CrearProducto(peso, tipoProd)
			if err != nil {
				fmt.Println(err)
			} else {
				flete.AgregarProducto(prod)
				var opcion string
				for !(opcion == "y" || opcion == "n") {
					fmt.Println("Desea agregar otro producto? Y or N")
					fmt.Scanln(&opcion)
				}
				if opcion == "n" {
					bucle = false
				}
			}
		}

	}
	envios := flete.CalcularEnvios()
	fmt.Println("Se deberan realizar ", envios, " envios")
}
