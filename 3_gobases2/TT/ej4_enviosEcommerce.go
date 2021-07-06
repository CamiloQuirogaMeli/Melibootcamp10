package main

import "fmt"

/*
Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un adicional al momento de ser enviado:
● Chico: Ningún adicional.
● Mediano: Requiere un %5 más de espacio
● Grande: Requiere un %20 más de espacio
● Frágil: Requiere un %75 más de espacio
● Especial: Sólo puede ser enviado con productos especiales

Se solicita que:
1. Los productos guarden el tamaño y tengan un método Tamaño Total que nos devuelva el espacio en cm3 que requerimos para ser enviado.
2. Y una estructura Flete que tenga los métodos:
- Agregar Producto: agregar producto al flete.
- Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede cargar un total de 10.000.000 cm3 por envío.
*/

const (
	prod_chico    = "chico"
	prod_mediano  = "mediano"
	prod_grande   = "grande"
	prod_fragil   = "fragil"
	prod_especial = "especial"
	LIMITE        = 10000000
)

type Flete struct {
	ProductosFlete []ProductoEnvios
}

type Pchico struct {
	tamano float64
}

func (p Pchico) tamanoTotal() float64 {
	return p.tamano
}

type Pmediano struct {
	tamano float64
}

func (p Pmediano) tamanoTotal() float64 {
	return p.tamano * 1.05
}

type Pgrande struct {
	tamano float64
}

func (p Pgrande) tamanoTotal() float64 {
	return p.tamano * 1.2
}

type Pfragil struct {
	tamano float64
}

func (p Pfragil) tamanoTotal() float64 {
	return p.tamano * 1.75
}

type Pespecial struct {
	tamano float64
}

func (p Pespecial) tamanoTotal() float64 {
	return p.tamano
}

type ProductoEnvios interface {
	tamanoTotal() float64
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

/*
func productoDetalle(p ProductoEnvios){
	fmt.Println(p)
	fmt.Println(p.agregarProducto())
	fmt.Println(p.calcularEnvio())

}
*/

func (f *Flete) agregarProducto(prod ...ProductoEnvios) {
	f.ProductosFlete = append(f.ProductosFlete, prod...)
}

func (f Flete) calcularEnvio() (cantEnviosOtros int, cantEnviosEspeciales int) {
	var especiales, otros []ProductoEnvios
	var tamanoEspeciales, tamanoOtros int = 0, 0
	for _, value := range f.ProductosFlete {
		if typeOf(value) == "main.Pespecial" {
			especiales = append(especiales, value)
		} else {
			otros = append(otros, value)
		}
	}
	for _, value := range otros {
		tamanoOtros += int(value.tamanoTotal())
	}
	for _, value := range especiales {
		tamanoEspeciales += int(value.tamanoTotal())
	}

	cantEnviosEspeciales = (tamanoEspeciales / LIMITE)
	if (tamanoEspeciales % 100000) > 0 {
		cantEnviosEspeciales = cantEnviosEspeciales + 1
	}

	cantEnviosOtros = (tamanoOtros / LIMITE)
	if (tamanoOtros % 100000) > 0 {
		cantEnviosOtros = cantEnviosOtros + 1
	}

	return

}

func fabrica(tipo string, tamano float64) ProductoEnvios {
	switch tipo {
	case "chico":
		return Pchico{tamano: tamano}
	case "mediano":
		return Pmediano{tamano: tamano}
	case "grande":
		return Pgrande{tamano: tamano}
	case "fragil":
		return Pfragil{tamano: tamano}
	case "especial":
		return Pespecial{tamano: tamano}
	}
	return nil
}
