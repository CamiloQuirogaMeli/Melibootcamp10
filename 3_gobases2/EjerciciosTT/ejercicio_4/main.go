package main

import "fmt"

/*
const (
	chico    = "chico"
	mediano  = "mediano"
	grande   = "grande"
	especial = "especial"
	fragil   = "fragil"
)
*/

type Producto struct {
	tipoProducto string
	tamanio      float64
}

type Flete struct {
	productos []Producto
}

func (p Producto) tamanioTotal() float64 {
	switch p.tipoProducto {
	case "chico":
		return p.tamanio
	case "mediano":
		return p.tamanio * 1.5
	case "grande":
		return p.tamanio * 1.2
	case "especial":
		return p.tamanio * 1.75
	case "fragil":
		return 0
	default:
		return 0
	}
}

func (f *Flete) agregarProducto(p Producto) {
	f.productos = append(f.productos, p)
}

func (f Flete) calcularEnvios() {
	var tamanioTotal float64 = 0
	// variable para controlar que no haya productos Especiales y Otros
	var cantProdEsp int = 0

	for _, p := range f.productos {
		if p.tipoProducto == "especial" {
			cantProdEsp++
		}
		tamanioTotal += p.tamanioTotal()
	}

	if cantProdEsp < len(f.productos) {
		fmt.Println("Productos Especiales y Otros - No es posible realizar el envío.")
		return
	}

	if tamanioTotal > 10000000 {
		fmt.Println("Exceso de Capacidad - No es posible realizar el envío.")
		return
	}

	fmt.Printf("Es posible realizar el envío. La capacidad total es de %.2f cm3", tamanioTotal)
}

func main() {
	flete := Flete{}
	p1 := Producto{tipoProducto: "chico", tamanio: 19878760}
	p2 := Producto{tipoProducto: "mediano", tamanio: 100}
	p3 := Producto{tipoProducto: "grande", tamanio: 10}
	p4 := Producto{tipoProducto: "grande", tamanio: 100}
	p5 := Producto{tipoProducto: "fragil", tamanio: 10}
	p6 := Producto{tipoProducto: "mediano", tamanio: 100}
	p7 := Producto{tipoProducto: "grande", tamanio: 87665}
	p8 := Producto{tipoProducto: "especial", tamanio: 3000}

	flete.agregarProducto(p1)
	flete.agregarProducto(p2)
	flete.agregarProducto(p3)
	flete.agregarProducto(p4)
	flete.agregarProducto(p5)
	flete.agregarProducto(p6)
	flete.agregarProducto(p7)
	flete.agregarProducto(p8)

	flete.calcularEnvios()
}
