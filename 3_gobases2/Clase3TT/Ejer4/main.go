package main

import (
	"fmt"
	"math"
	"reflect"
)

const (
	chico    = "little"
	mediano  = "medium"
	grande   = "big"
	especial = "special"
	fragil   = "fagile"
)

type Product interface {
	TamanoTotal() float64
}
type little struct {
	size float64
}

func (l little) TamanoTotal() float64 {
	return l.size
}

type medium struct {
	size float64
}

func (m medium) TamanoTotal() float64 {
	return m.size + m.size*0.05
}

type big struct {
	size float64
}

func (b big) TamanoTotal() float64 {
	return b.size + b.size*0.2
}

type fragile struct {
	size float64
}

func (f fragile) TamanoTotal() float64 {
	return f.size + f.size*0.75
}

type special struct {
	size float64
}

func (s special) TamanoTotal() float64 {
	return s.size
}

type flete struct {
	products []Product
}

func (f *flete) AgregarProducto(p Product) {
	f.products = append(f.products, p)
}

func (f *flete) CalcularEnvios() float64 {
	regular, special := 0.0, 0.0
	for _, product := range f.products {
		if reflect.TypeOf(product).Name() == especial {
			special += product.TamanoTotal()
		} else {
			regular += product.TamanoTotal()
		}
	}
	return math.Ceil(regular/10000000.0) + math.Ceil(special/10000000)
}

func main() {
	flete := flete{}
	flete.AgregarProducto(little{10000000 - 1})
	flete.AgregarProducto(little{2})
	flete.AgregarProducto(special{10000001})
	fmt.Println("El numero de envios necesarios es de", flete.CalcularEnvios())
}
