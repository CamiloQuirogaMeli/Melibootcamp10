package main

import (
	"fmt"
)

const (
	S   = "small"
	M   = "medium"
	B   = "big"
	X   = "special"
	F   = "fragile"
	MAX = 10000.0
)

type product struct {
	name       string
	size       string
	cubicMeter float64
}

type freight struct {
	load []product
}

var (
	spaceInfo = map[string]float64{
		S: 0.0,
		M: 0.05,
		B: 0.2,
		X: 0.0,
		F: 0.75,
	}
)

func main() {

	flete, productos := startUp()

	for _, value := range productos {
		//fmt.Println("Nombre: ", value.name, " | Size:", value.size, " | TotalSpace:", value.getTotalSpace())
		flete.addProduct(value)
	}

	shippingQuantity := flete.getShippingQuantity()

	fmt.Println(shippingQuantity)
}

func startUp() (freight, []product) {

	products := []product{
		{"lampara", "small", 56.33},
		{"alacena", "medium", 2345.2},
		{"ropero", "big", 523423.4},
		{"pc", "fragile", 2345.3},
		{"camioneta", "special", 20000.1},
	}

	return freight{}, products
}

func (p product) getTotalSpace() float64 {
	return p.cubicMeter + (p.cubicMeter * spaceInfo[p.size])
}

func (f *freight) addProduct(p product) {
	f.load = append(f.load, p)
}

func (f freight) getShippingQuantity() int {

	normalItemUsedSpace := 0.0
	specialItemUsedSpace := 0.0

	for _, value := range f.load {

		if value.size != X {
			normalItemUsedSpace += value.getTotalSpace()
		} else {
			specialItemUsedSpace += value.getTotalSpace()
		}
	}

	normItemResult := getItemShipQuant(normalItemUsedSpace)
	specialItemResult := getItemShipQuant(specialItemUsedSpace)

	return normItemResult + specialItemResult
}

func getItemShipQuant(q float64) int {
	result := int(q) / MAX

	if result == 0 {
		result = 1
	}

	return result
}

/*
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
