/*
   Ejercicio 4 - Envios
   Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de
   productos:
   La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
   Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto
   requiere un adicional al momento de ser enviado:
   ● Chico: Ningún adicional.
   ● Mediano: Requiere un %5 más de espacio
   ● Grande: Requiere un %20 más de espacio
   ● Frágil: Requiere un %75 más de espacio
   ● Especial: Sólo puede ser enviado con productos especiales

   Se solicita que:
   1. Los productos guarden el tamaño y tengan un método Tamaño Total que nos
   devuelva el espacio en cm3 que requerimos para ser enviado.
   2. Y una estructura Flete que tenga los métodos:
   - Agregar Producto: agregar producto al flete.
   - Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo
   puede cargar un total de 10.000.000 cm3 por envío.
*/

package main

import "fmt"

const (
	CHICO    = "chico"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
	FRAGIL   = "fragil"
	ESPECIAL = "especial"
)

func main() {
	prod1 := factory(CHICO, 111110)
	prod2 := factory(MEDIANO, 2196662)
	prod3 := factory(CHICO, 781817)
	prod4 := factory(FRAGIL, 333245)
	prod5 := factory(ESPECIAL, 11334590)
	prod6 := factory(ESPECIAL, 111000000)

	freight := Freight{}

	freight.addProd(prod1, prod2, prod3, prod4, prod5, prod6)

	cantOtros, cantEsp := freight.shipping()

	fmt.Println("La cantidad de viajes especiales será de: ", cantEsp)
	fmt.Println("La cantidad de viajes de productos no especiales será de: ", cantOtros)
	fmt.Println("La cantidad total de viajes es de: ", cantEsp+cantOtros)

}

func (f Freight) shipping() (cantShippingSpecial int, cantShippingOth int) {
	var special, other []Producto
	var specialCM3, otherCM3 int = 0, 0

	for _, value := range f.Productos {
		if typeof(value) == "main.ProdEspecial" {
			special = append(special, value)
		} else {
			other = append(other, value)
		}
	}

	for _, value := range other {
		otherCM3 += int(value.totalSize())
	}

	for _, value := range special {
		specialCM3 += int(value.totalSize())
	}

	cantShippingOth = (otherCM3 / 10000000)
	if (otherCM3 % 1000000) > 0 {
		cantShippingOth += 1
	}

	cantShippingSpecial = (specialCM3 / 10000000)
	if (specialCM3 % 1000000) > 0 {
		cantShippingSpecial += 1
	}

	return
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
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

type Freight struct {
	Productos []Producto
}

func (f *Freight) addProd(prod ...Producto) {
	f.Productos = append(f.Productos, prod...)
}

type Producto interface {
	totalSize() float64
}

type ProdChico struct {
	cm3 float64
}

func (p ProdChico) totalSize() float64 {
	return p.cm3
}

type ProdMed struct {
	cm3 float64
}

func (p ProdMed) totalSize() float64 {
	return p.cm3 * 1.05
}

type ProdGrande struct {
	cm3 float64
}

func (p ProdGrande) totalSize() float64 {
	return p.cm3 * 1.2
}

type ProdFragil struct {
	cm3 float64
}

func (p ProdFragil) totalSize() float64 {
	return p.cm3 * 1.75
}

type ProdEspecial struct {
	cm3 float64
}

func (p ProdEspecial) totalSize() float64 {
	return p.cm3
}
