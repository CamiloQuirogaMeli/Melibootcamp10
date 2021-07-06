package main

import "fmt"

const (
	CHICO    = "CHICO"
	MEDIANO  = "MEDIANO"
	GRANDE   = "GRANDE"
	ESPECIAL = "ESPECIAL"
	FRAGIL   = "FRAGIL"
)

type Producto interface {
	TamanoTotal() float32
	isSpecial() bool
}

type Chico struct {
	Tamano   float32
	especial bool
}

type Mediano struct {
	Tamano   float32
	especial bool
}

type Grande struct {
	Tamano   float32
	especial bool
}

type Fragil struct {
	Tamano   float32
	especial bool
}

type Especial struct {
	Tamano   float32
	especial bool
}

func (c Chico) TamanoTotal() float32 {
	return c.Tamano
}

func (c Mediano) TamanoTotal() float32 {
	return c.Tamano + (c.Tamano * 0.05)
}
func (c Grande) TamanoTotal() float32 {
	return c.Tamano + (c.Tamano * 0.2)
}

func (c Fragil) TamanoTotal() float32 {
	return c.Tamano + (c.Tamano * 0.7)
}
func (c Especial) TamanoTotal() float32 {
	return c.Tamano
}

func (c Chico) isSpecial() bool {
	return c.especial
}

func (c Mediano) isSpecial() bool {
	return c.especial
}
func (c Grande) isSpecial() bool {
	return c.especial
}

func (c Fragil) isSpecial() bool {
	return c.especial
}
func (c Especial) isSpecial() bool {
	return c.especial
}

func factory(productType string, tamano float32) Producto {
	switch productType {
	case CHICO:
		return Chico{tamano, false}
	case MEDIANO:
		return Mediano{tamano, false}
	case GRANDE:
		return Grande{tamano, false}
	case FRAGIL:
		return Fragil{tamano, false}
	case ESPECIAL:
		return Especial{tamano, true}
	}
	return nil
}

type Flete struct {
	productosComunes    []Producto
	productosEspeciales []Producto
}

func (flete *Flete) AgregarProducto(producto Producto) {

	if !producto.isSpecial() {
		flete.productosComunes = append(flete.productosComunes, producto)
	} else {
		flete.productosEspeciales = append(flete.productosEspeciales, producto)
	}

}

func (flete *Flete) CalcularEnvios() {

	var (
		comunesTotal         float32
		cantEnviosComunes    int
		especialesTotal      float32
		cantEnviosEspeciales int
	)

	for i, producto := range flete.productosComunes {
		comunesTotal = comunesTotal + producto.TamanoTotal()

		if comunesTotal > 10000000 {
			fmt.Println("se realizara un envio de comunes con una capacidad de ", comunesTotal)
			comunesTotal = 0
			cantEnviosComunes++
			continue
		}
		if i == len(flete.productosComunes)-1 {
			cantEnviosComunes++
			fmt.Println("se realizara un envio de comunes con una capacidad de ", comunesTotal)
		}

	}

	for i, producto := range flete.productosEspeciales {
		especialesTotal = especialesTotal + producto.TamanoTotal()
		if especialesTotal > 10000000 {
			fmt.Println("se realizara un envio de especiales con una capacidad de ", especialesTotal)
			cantEnviosEspeciales++
			especialesTotal = 0
			continue
		}
		if i == len(flete.productosEspeciales)-1 {
			cantEnviosEspeciales++
			fmt.Println("se realizara un envio de especiales con una capacidad de ", especialesTotal)
		}
	}

	fmt.Printf("En total se realizaran %d envios comunes y %d envios especiales.", cantEnviosComunes, cantEnviosEspeciales)

}

func main() {

	flete := Flete{}

	flete.AgregarProducto(factory(CHICO, 90000000))
	flete.AgregarProducto(factory(MEDIANO, 12000000))
	flete.AgregarProducto(factory(GRANDE, 5000))
	flete.AgregarProducto(factory(ESPECIAL, 9000))
	flete.CalcularEnvios()

}
