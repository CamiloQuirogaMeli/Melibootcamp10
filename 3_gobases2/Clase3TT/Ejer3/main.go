package main

import "fmt"

const (
	chico   = "chico"
	mediano = "mediano"
	grande  = "grande"
)

const (
	shopA = "tiendaUno"
	shopB = "tiendaDos"
)

type Ecommerce interface {
	precio() float64
	Envio() string
}

type tiendaUno struct {
	nombre    string
	products  []Product
	direccion string
}

func (shopA tiendaUno) precio() float64 {
	total := 0.0
	for _, product := range shopA.products {
		total += product.precio()
	}
	return total
}

func (shopA tiendaUno) Envio() string {
	return shopA.direccion
}

type tiendaDos struct {
	nombre    string
	products  []Product
	direccion string
}

func (shopB tiendaDos) precio() float64 {
	total := 0.0
	for _, product := range shopB.products {
		total += product.precio()
	}
	return total
}

func (shopB tiendaDos) Envio() string {
	return shopB.direccion
}

func nuevaTienda(shopType string, name string, products []Product, address string) Ecommerce {
	switch shopType {
	case shopA:
		return tiendaUno{name, products, address}
	case shopB:
		return tiendaUno{name, products, address}
	default:
		return nil
	}
}

func factory(itemType string, values ...float64) Product {
	switch itemType {
	case chico:
		return little{values[0]}
	case mediano:
		return medium{values[0], values[1]}
	case grande:
		return big{values[0], values[1], values[2]}
	default:
		return nil
	}
}

type Product interface {
	precio() float64
}

type little struct {
	p float64
}

func (l little) precio() float64 {
	return l.p
}

type medium struct {
	p      float64
	charge float64
}

func (m medium) precio() float64 {
	return m.p + m.p*m.charge
}

type big struct {
	p      float64
	charge float64
	feed   float64
}

func (b big) precio() float64 {
	return b.p + b.p*b.charge + b.feed
}

func main() {
	pTiedaA := []Product{little{1000}, medium{12341, 0.10}, big{1000, 0.1, 500}}
	pTiedaB := []Product{big{1000, 0.1, 100}, big{1000, 0.234, 100}, big{1000, 0.1, 100}}

	t1 := nuevaTienda(shopA, "Tienda A", pTiedaA, "calle 46")
	t2 := nuevaTienda(shopB, "Tienda 1", pTiedaB, "calle 64")

	fmt.Println(t1.precio())
	fmt.Println(t1.Envio())
	fmt.Println(t2.precio())
	fmt.Println(t2.Envio())
}
