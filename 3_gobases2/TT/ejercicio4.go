package main

import "fmt"

const (
	chico    = "chico"
	mediano  = "mediano"
	grande   = "grande"
	fragil   = "fragil"
	especial = "especial"
)

type Flete struct {
	productos []Producto
}

type Producto struct {
	tipo   string
	tamano float64
}

func (p Producto) ObtenerEspacio() float64 {
	switch p.tipo {
	case chico:
		return p.tamano
	case mediano:
		return p.tamano * 1.05
	case grande:
		return p.tamano * 1.2
	case fragil:
		return p.tamano * 1.75
	default:
		return p.tamano
	}
}

func (f *Flete) AgregarProducto(prod Producto) {
	if prod.tipo != especial {
		f.productos = append(f.productos, prod)
		return
	}
	for _, product := range f.productos {
		if product.tipo != especial {
			return
		}
	}
	f.productos = append(f.productos, prod)

}

func (f Flete) CalcularEnvios() int {
	var espacioTotal float64

	for _, prod := range f.productos {
		espacioTotal += prod.tamano
	}
	var res int = int(espacioTotal) % 10000000

	if res == 0 {
		return int(espacioTotal) / 10000000
	} else {
		return (int(espacioTotal) / 10000000) + 1
	}

}

func main() {

	flete := Flete{}
	flete2 := Flete{}

	prod := Producto{tipo: chico, tamano: 5000000}
	prod2 := Producto{tipo: mediano, tamano: 10000000}
	prod3 := Producto{tipo: grande, tamano: 50000000}
	prod4 := Producto{tipo: especial, tamano: 4000000}
	prod5 := Producto{tipo: especial, tamano: 6000000}

	flete.AgregarProducto(prod)
	flete.AgregarProducto(prod2)
	flete.AgregarProducto(prod3)
	flete.AgregarProducto(prod4) // no se agregan productos especuales a este flete

	// Productos especiales
	flete2.AgregarProducto(prod4)
	flete2.AgregarProducto(prod5)

	// flete 1
	fmt.Println("Flete 1:")
	for i, prod := range flete.productos {
		fmt.Printf("producto %d - espacio: %.2f cm3 \n", i+1, prod.tamano)
	}
	fmt.Println("No Envios:", flete.CalcularEnvios())
	fmt.Println("\nFlete 2:")
	// flete 2
	for i, prod := range flete2.productos {
		fmt.Printf("producto %d - espacio: %.2f cm3 \n", i+1, prod.tamano)
	}
	fmt.Println("No Envios:", flete2.CalcularEnvios())
}
