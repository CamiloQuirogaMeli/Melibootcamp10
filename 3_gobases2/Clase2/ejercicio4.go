package main

import (
	"fmt"
)

type Flete struct {
	name      string
	productos []Producto
}

func (f *Flete) AgregarProducto(producto Producto) {

	f.productos = append(f.productos, producto)
	fmt.Println("El producto", producto.name, "de tipo", producto.tipo, "ha sido agregado")

}

func (f *Flete) CalcularEnvios() {
	var totalEnvios float64
	for _, producto := range f.productos {
		totalEnvios += producto.tamano
	}
	fmt.Println("Para el flete", f.name)
	fmt.Println("Se tienen", len(f.productos), "productos")
	fmt.Println("Que suman un total de", totalEnvios, "centimetros cuadrados")
}

type Producto struct {
	name   string
	tipo   string
	tamano float64
}

func (p *Producto) TamanoTotal() {
	var espacio float64
	switch p.tipo {
	case "c", "C":
		espacio = p.tamano
		fmt.Println("No se requiere espacio adicional")
	case "m", "M":
		espacio = p.tamano * 1.05
		fmt.Println("Se requiere un 5% mas de espacio")
	case "g", "G":
		espacio = p.tamano * 1.2
		fmt.Println("Se requiere un 20% mas de espacio")
	case "f", "F":
		espacio = p.tamano * 1.75
		fmt.Println("Se requiere un 75% mas de espacio")
	case "e", "E":
		espacio = 0
		fmt.Println("Sólo puede ser enviado con productos especiales")
	}
	if p.tipo != "e" && p.tipo != "E" {
		fmt.Println("el espacio en cm3 que requerimos para ser enviado es de", espacio)

	}
}

func main() {
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
	flete := Flete{}
	salir := false
	var accion int
	productos := make([]Producto, 0)

	fmt.Println("Digita el nombre del flete")
	fmt.Scanln(&flete.name)

	for !salir {
		fmt.Println("Digita la accion a realizar:")
		fmt.Println("	1: AgregarProducto")
		fmt.Println("	2: CalcularEnvios")
		fmt.Println("	3: TamanoTotal")
		fmt.Println("	4: Salir")
		fmt.Scanln(&accion)
		switch accion {
		case 1:
			productos = append(productos, Producto{})
			fmt.Println("Digita el nombre del producto")
			fmt.Scanln(&productos[len(productos)-1].name)
			fmt.Println("Digita el tipo del producto")
			fmt.Scanln(&productos[len(productos)-1].tipo)
			fmt.Println("Digita el tamaño del producto")
			fmt.Scanln(&productos[len(productos)-1].tamano)

			flete.AgregarProducto(productos[len(productos)-1])
		case 2:

			flete.CalcularEnvios()

		case 3:
			fmt.Println("Los tamaños de cada producto son los siguientes")
			for i, producto := range flete.productos {
				fmt.Println("Para el producto", i+1)
				producto.TamanoTotal()
			}

		case 4:
			salir = true
		default:
			fmt.Println("Opción no valida")
		}

	}

}
