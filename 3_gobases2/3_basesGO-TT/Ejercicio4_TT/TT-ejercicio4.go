package main

import "fmt"

/*
Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de
productos:
La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto
requiere un adicional al momento de ser enviado:
Chico: Ningún adicional.
Mediano: Requiere un %5 más de espacio
Grande: Requiere un %20 más de espacio
Frágil: Requiere un %75 más de espacio
Especial: Sólo puede ser enviado con productos especiales

Para ello requerimos que los productos guarden el tamaño y tengan un metodo TamanoTotal
que nos devuelva el el espacio en cm3 que requerimos para ser enviado.

Y una estructura Flete que tenga los métodos:
- AgregarProducto: agregar producto al flete
- CalcularEnvios: calcula la cantidad de envíos que debe realizar sabiendo que solo
puede cargar un total de 10.000.000 cm3 por envío
*/

const (
	CHICO    = "Chico"
	MEDIANO  = "Mediano"
	GRANDE   = "Grande"
	FRAGIL   = "Fagil"
	ESPECIAL = "Especial"
)

type producto struct {
	nombre     string
	tamCm3     float64
	tipo       string
	esEspecial bool
}

func (p producto) tamanoTotal() float64 {

	tam := p.tamCm3

	switch p.tipo {
	case CHICO:
		return tam
	case MEDIANO:
		return tam * 1.05
	case GRANDE:
		return tam * 1.2
	case FRAGIL:
		return tam * 1.75
	case ESPECIAL:
		return tam
	default:
		return 0
	}

}

type flete struct {
	idFlete   int
	productos []producto
}

func (f *flete) agregarProducto(p producto) {
	f.productos = append(f.productos, p)
}

func (f flete) calcularEnvios() (int, int) {

	cantFletes := 1
	cantProductosEsp := 0
	produ := f.productos
	cantCm3 := 0.0

	for _, prod := range produ {

		switch {
		/*Si es especial, va a envio especial*/
		case prod.esEspecial:
			cantProductosEsp++
			/*Si el producto entra en el envio, lo sumo*/
		case (cantCm3 + prod.tamanoTotal()) <= 10000000.0:
			cantCm3 += prod.tamanoTotal()
			/*Si el envio esta lleno, empezamos otro*/
		default:
			cantFletes++
			cantCm3 = prod.tamanoTotal()
		}
	}
	return cantFletes, cantProductosEsp
}

func main() {

	p1 := producto{"bicicleta", 1000.0, "Grande", false}
	fmt.Println(p1)
	p2 := producto{"anillo", 1000.0, "Chico", false}
	fmt.Println(p2)
	p3 := producto{"espejo", 500.0, "Especial", true}
	fmt.Println(p3)
	p4 := producto{"cocina", 1000.0, "Mediano", false}
	fmt.Println(p4)

	p := []producto{p1, p2, p3}

	f1 := flete{1, p}
	f1.agregarProducto(p4)
	cFletes, cEsp := f1.calcularEnvios()
	fmt.Println("Cantidad de fletes: ", cFletes)
	fmt.Println("Cantidad de prodEspeciales: ", cEsp)

}
