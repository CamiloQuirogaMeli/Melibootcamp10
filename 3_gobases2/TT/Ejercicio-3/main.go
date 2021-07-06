package main

import "fmt"

type tiendaTodoPC struct {
	nombre string
	rubro  string
	insta  string
	precio float64
}
type tiendaCoverClothes struct {
	nombre string
	rubro  string
	web    string
	precio float64
}

type Ecommerce interface {
	Precio(tamanioProd string) float64
	Envio(address string) string
}

func (e tiendaCoverClothes) Precio(tamanioProd string) float64 {
	switch tamanioProd {
	case "small":
		return e.precio
	case "medium":
		return e.precio * 1.03
	case "big":
		return (e.precio * 1.06) + 2500
	}
	return e.precio
}
func (e tiendaCoverClothes) Envio(address string) string {
	sent := "El envio sera enviado a la direccion" + address
	return sent
}

func (e tiendaTodoPC) Precio(tamanioProd string) float64 {
	switch tamanioProd {
	case "small":
		return e.precio
	case "medium":
		return e.precio * 1.03
	case "big":
		return (e.precio * 1.06) + 2500
	}
	return e.precio
}

func (e tiendaTodoPC) Envio(address string) string {
	sent := "El envio sera enviado a la direccion" + address
	return sent
}
func nuevaTienda(productType string) Ecommerce {
	if productType == "tiendaTodoPC" {
		return tiendaTodoPC{nombre: "memoriaRam", rubro: "computacion", insta: "instagram.com/todopc", precio: 9000}
	}
	if productType == "tiendaCoverClothes" {
		return tiendaCoverClothes{nombre: "Buzo", rubro: "Abrigos", web: "www.CoverClothes.com.ar", precio: 3500}
	}
	return nil
}

func main() {
	fmt.Println("Ejercicio 3 TT")
	tiendaPC := nuevaTienda("tiendaTodoPC")
	pricePC := tiendaPC.Precio("small")
	envioPC := tiendaPC.Envio("Av. Salvador Gallardo Nro 1328")
	fmt.Println(tiendaPC)
	fmt.Println("El precio del producto sera:", pricePC)
	fmt.Println(envioPC)

	fmt.Println()

	tiendaRopa := nuevaTienda("tiendaCoverClothes")
	precioRopa := tiendaRopa.Precio("medium")
	envioRopa := tiendaRopa.Envio("Chacabuco 2322")
	fmt.Println(tiendaRopa)
	fmt.Println("El precio del producto sera", precioRopa)
	fmt.Println(envioRopa)

}
