package main

import "fmt"

const (
	CHICO    = "chico"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
	FRAGIL   = "fragil"
	ESPECIAL = "especial"
)

type Producto struct {
	tamaño float64
	tipo   string
}

type Flete struct {
	productos     []Producto
	tieneEspecial bool
}

func (p Producto) tamañoTotal(prodType string) float64 {
	switch prodType {
	case CHICO:
		return p.tamaño
	case ESPECIAL:
		return p.tamaño
	case MEDIANO:
		return p.tamaño * 1.05
	case GRANDE:
		return p.tamaño * 1.2
	case FRAGIL:
		return p.tamaño * 1.75
	}
	return 0
}

func (f *Flete) agregarProducto(p Producto) {
	if p.tipo == ESPECIAL {
		if len(f.productos) == 0 {
			f.productos = append(f.productos, p)
			f.tieneEspecial = true
		} else {
			if f.tieneEspecial {
				f.productos = append(f.productos, p)
			} else {
				fmt.Println("No se puede agregar un producto especial a este flete")
			}
		}
	} else {
		if !f.tieneEspecial {
			f.productos = append(f.productos, p)
		} else {
			fmt.Println("No se puede agregar un producto NO especial a este flete")
		}

	}
}

func (f Flete) calcularEnvios() int {
	capacidadMax := 0.0
	for _, prod := range f.productos {
		capacidadMax += prod.tamañoTotal(prod.tipo)
	}
	return (int(capacidadMax) / 10000000)
}

func main() {
	p1 := Producto{1000, GRANDE}
	p2 := Producto{50, MEDIANO}
	p3 := Producto{10000001, CHICO}

	f1 := Flete{}

	f1.agregarProducto(p1)
	f1.agregarProducto(p2)
	f1.agregarProducto(p3)
	fmt.Println(f1.productos)

	fmt.Println("La cantidad de envios a hacer es de ", f1.calcularEnvios()+1)
}
