package main

import "fmt"

func main() {
	tiendaUno := nuevaTienda("tiendauno")
	precioUno := tiendaUno.Precio("mediano")
	envioUno := tiendaUno.Envio("Av independencia 333, Sunchales")
	tiendaDos := nuevaTienda("tiendados")
	precioDos := tiendaUno.Precio("grande")
	envioDos := tiendaUno.Envio("San Martin 400, Sunchales")

	fmt.Println("================================")
	fmt.Println(tiendaUno)
	fmt.Println("El precio final del producto es:", precioUno)
	fmt.Println(envioUno)

	fmt.Println("================================")
	fmt.Println(tiendaDos)
	fmt.Println("El precio final del producto es:", precioDos)
	fmt.Println(envioDos)
}

type tiendaUno struct {
	direccionTienda string
	precio          float64
}

type tiendaDos struct {
	direccionTienda string
	precio          float64
	paginaWeb       string
}

type Ecommerce interface {
	Precio(tamaño string) float64
	Envio(direccion string) string
}

func nuevaTienda(nombreTienda string) Ecommerce {
	if nombreTienda == "tiendauno" {
		return tiendaUno{direccionTienda: "Calle falsa 123", precio: 2500.00}
	}
	if nombreTienda == "tiendados" {
		return tiendaDos{direccionTienda: "Calle falsa 123456", precio: 2300.00, paginaWeb: "tiendaDos.com"}
	}
	return nil
}

func (tuno tiendaUno) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return tuno.precio
	case "mediano":
		precioFinal := (tuno.precio * 1.3)
		return precioFinal
	case "grande":
		precioFinal := ((tuno.precio * 1.6) + 2500)
		return precioFinal
	default:
		return 0
	}
}

func (tuno tiendaUno) Envio(direccion string) string {
	return fmt.Sprintf("El envio se realizara a: %s", direccion)
}

func (tdos tiendaDos) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return tdos.precio
	case "mediano":
		precioFinal := (tdos.precio * 1.3)
		return precioFinal
	case "grande":
		precioFinal := ((tdos.precio * 1.6) + 2500)
		return precioFinal
	default:
		return 0
	}
}

func (tdos tiendaDos) Envio(direccion string) string {
	return fmt.Sprintf("El envio se realizara a: %s", direccion)
}
