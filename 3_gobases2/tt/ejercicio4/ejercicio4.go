package main

import "fmt"

const (
	smallType   = 1
	mediumType  = 2
	bigType     = 3
	frailType   = 4
	specialType = 5
	capacidad   = 10000000
)

type Freight struct {
	Prod []Product
}

type Product struct {
	typ  int16
	size float64
}

func (p Product) tamanioTotal() float64 {
	switch p.typ {
	case 1:
		return p.size
	case 2:
		return p.size * 1.5
	case 3:
		return p.size * 1.2
	case 4:
		return p.size * 1.75
	case 5:
		return p.size
	default:
		return 0
	}
}

func (f *Freight) agregarProducto(p Product) {
	f.Prod = append(f.Prod, p)
}

func (f Freight) calcularEnvio() {
	var tamanioTotal float64 = 0
	var tamanioProdEsp float64 = 0

	for _, p := range f.Prod {
		if p.typ == 5 {
			tamanioProdEsp += p.tamanioTotal()
		} else {
			tamanioTotal += p.tamanioTotal()
		}
	}
	cantEnvios := int32(tamanioTotal) / int32(capacidad)
	if int32(tamanioTotal)%int32(capacidad) > 0 {
		cantEnvios++
	}
	cantEnviosEsp := int32(tamanioProdEsp) / int32(capacidad)
	if int32(tamanioProdEsp)%int32(capacidad) > 0 {
		cantEnviosEsp++
	}

	fmt.Printf("Se deben realizar %d envíos de Productos Especiales y Otros\n", cantEnviosEsp)
	fmt.Printf("Se deben realizar %d envíos de Productos Chicos, Medianos, Grandes y Frágiles\n", cantEnvios)
}

func main() {
	flete := Freight{}
	p1 := Product{typ: 1, size: 10000000}
	p2 := Product{typ: 2, size: 100}
	p3 := Product{typ: 3, size: 10}
	p4 := Product{typ: 3, size: 100}
	p5 := Product{typ: 4, size: 10}
	p6 := Product{typ: 2, size: 100}
	p7 := Product{typ: 3, size: 30000000}
	p8 := Product{typ: 5, size: 3000000}
	p9 := Product{typ: 5, size: 7000010}

	flete.agregarProducto(p1)
	flete.agregarProducto(p2)
	flete.agregarProducto(p3)
	flete.agregarProducto(p4)
	flete.agregarProducto(p5)
	flete.agregarProducto(p6)
	flete.agregarProducto(p7)
	flete.agregarProducto(p8)
	flete.agregarProducto(p9)

	flete.calcularEnvio()
}
