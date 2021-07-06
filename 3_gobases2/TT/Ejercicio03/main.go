package main

import "fmt"

type Ecommerce interface {
	Precio(tamano string) float64
	Envio(domicilio string) string
}

type TiendaCamisas struct {
	Nombre, Domicilio string
}

type TiendaDeportes struct {
	Nombre string
}

func (t TiendaCamisas) Precio(tamano string) float64 {
	precioProducto := map[string]float64{"Chico": 150, "Mediano": 750, "Grande": 1250}
	precio := precioProducto[tamano]

	if tamano == "Mediano" {
		precio = precio + (precio * 3 / 100)
	} else if tamano == "Grande" {
		precio = precio + (precio * 6 / 100) + 2500
	}
	return precio
}

func (t TiendaCamisas) Envio(domicilio string) string {
	return domicilio
}

func (t TiendaDeportes) Precio(tamano string) float64 {
	precioProducto := map[string]float64{"Chico": 100, "Mediano": 450, "Grande": 750}
	precio := precioProducto[tamano]

	if tamano == "Mediano" {
		precio = precio + (precio * 3 / 100)
	} else if tamano == "Grande" {
		precio = precio + (precio * 6 / 100) + 2500
	}
	return precio
}

func (t TiendaDeportes) Envio(domicilio string) string {
	return domicilio
}

func nuevaTienda(tipoDeProducto string) Ecommerce {
	switch tipoDeProducto {
	case "tiendaCamisas":
		return TiendaCamisas{Nombre: "Tienda 1"}
	case "tiendaDeportes":
		return TiendaDeportes{Nombre: "Tienda 2"}
	}
	return nil
}

func main() {
	const (
		camisas  = "tiendaCamisas"
		deportes = "tiendaDeportes"
	)
	t1 := nuevaTienda(camisas)
	fmt.Println(t1.Precio("Mediano"))
	fmt.Println(t1.Envio("Domicilio de entrega"))

	t2 := nuevaTienda(deportes)
	fmt.Println(t2.Envio("Mariano Acha 1234"))
	fmt.Println(t2.Precio("Mediano"))
}
