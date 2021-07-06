package main

import "fmt"

const (
	CHICO    = "CHICO"
	MEDIANO  = "MEDIANO"
	GRANDE   = "GRANDE"
	FRAGIL   = "FRAGIL"
	ESPECIAL = "ESPECIAL"
)

type Flete struct {
	Productos []Producto
}

func (f *Flete) agregarProducto(prod ...Producto) {
	f.Productos = append(f.Productos, prod...)
}

func (f Flete) calcularEnvios() (cantEnviosOtros int, cantEnviosEspeciales int) {
	var especiales, otros []Producto
	var cm3Esepciales, cm3Otros int = 0, 0

	for _, value := range f.Productos {
		if typeof(value) == "main.ProdEspecial" {
			especiales = append(especiales, value)
		} else {
			otros = append(otros, value)
		}
	}

	for _, value := range otros {
		cm3Otros += int(value.tamanioTotal())
	}

	for _, value := range especiales {
		cm3Esepciales += int(value.tamanioTotal())
	}

	cantEnviosEspeciales = (cm3Esepciales / 10000000)
	if (cm3Esepciales % 1000000) > 0 {
		cantEnviosEspeciales = cantEnviosEspeciales + 1
	}

	cantEnviosOtros = (cm3Otros / 10000000)
	if (cm3Otros % 1000000) > 0 {
		cantEnviosOtros = cantEnviosOtros + 1
	}

	return
}

func factory(tipo string, cm float64) Producto {
	switch tipo {
	case CHICO:
		return ProdChico{cm}
	case MEDIANO:
		return ProdMed{cm}
	case GRANDE:
		return ProdGrande{cm}
	case FRAGIL:
		return ProdFragil{cm}
	case ESPECIAL:
		return ProdEspecial{cm}
	}
	return nil
}

type Producto interface {
	tamanioTotal() float64
}

type ProdChico struct {
	cm3 float64
}

func (p ProdChico) tamanioTotal() float64 {
	return p.cm3
}

type ProdMed struct {
	cm3 float64
}

func (p ProdMed) tamanioTotal() float64 {
	return p.cm3 * 1.05
}

type ProdGrande struct {
	cm3 float64
}

func (p ProdGrande) tamanioTotal() float64 {
	return p.cm3 * 1.2
}

type ProdFragil struct {
	cm3 float64
}

func (p ProdFragil) tamanioTotal() float64 {
	return p.cm3 * 1.75
}

type ProdEspecial struct {
	cm3 float64
}

func (p ProdEspecial) tamanioTotal() float64 {
	return p.cm3
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func main() {
	p1 := factory(CHICO, 100000)
	p2 := factory(MEDIANO, 250000)
	p3 := factory(GRANDE, 3245768)
	p4 := factory(FRAGIL, 786754)
	p5 := factory(ESPECIAL, 5100000)
	p6 := factory(ESPECIAL, 5000000)
	p7 := factory(ESPECIAL, 500)
	p8 := factory(GRANDE, 585858)
	p9 := factory(GRANDE, 123456789)

	f1 := Flete{}

	f1.agregarProducto(p1, p2, p3, p4, p5, p6, p7, p8, p9)

	cantOtros, cantEsp := f1.calcularEnvios()

	fmt.Println("La cantidad de viajes especiales será de: ", cantEsp)
	fmt.Println("La cantidad de viajes de productos no especiales será de: ", cantOtros)
	fmt.Println("La cantidad total de viajes es de: ", cantEsp+cantOtros)

}
