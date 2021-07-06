package main

import (
	"errors"
	"fmt"
)

const (
	PEQUENIO = "pequeño"
	MEDIANO  = "mediano"
	GRANDE   = "grande"
)

type Ecommerce interface {
	Precio(float64) float64
	Envio(string) string
}

type ProductoChico struct {
	precio float64
}

func (p ProductoChico) Precio(cantidad float64) float64 {
	return p.precio * cantidad
}

func (p ProductoChico) Envio(direccion string) string {
	return direccion
}

type ProductoMediano struct {
	precio       float64
	costoAlmacen float64
}

func (p ProductoMediano) Precio(cantidad float64) float64 {
	return cantidad * (p.precio * (1 + p.costoAlmacen))
}

func (p ProductoMediano) Envio(direccion string) string {
	return direccion
}

type ProductoGrande struct {
	precio       float64
	costoAlmacen float64
	costoEnvio   float64
}

func (p ProductoGrande) Precio(cantidad float64) float64 {
	return cantidad * (p.precio*(1+p.costoAlmacen) + p.costoEnvio)
}

func (p ProductoGrande) Envio(direccion string) string {
	return direccion
}

func nuevoProducto(tipoProducto string, precio float64) (Ecommerce, error) {
	switch tipoProducto {
	case PEQUENIO:
		return ProductoChico{precio}, nil
	case MEDIANO:
		return ProductoMediano{precio, 0.03}, nil
	case GRANDE:
		return ProductoGrande{precio, 0.06, 2500}, nil
	default:
		return nil, errors.New("No existe ese tipo de producto")
	}
}

func main() {
	producto, _ := nuevoProducto(PEQUENIO, 1200.50)
	fmt.Println("El precio es:", producto.Precio(1))
	fmt.Println("La dirección es:", producto.Envio("Mexico"))
	producto, _ = nuevoProducto(MEDIANO, 1000.0)
	fmt.Println("El precio es:", producto.Precio(1))
	fmt.Println("La dirección es:", producto.Envio("Mexico"))
	producto, _ = nuevoProducto(GRANDE, 100.0)
	fmt.Println("El precio es:", producto.Precio(1))
	fmt.Println("La dirección es:", producto.Envio("Mexico"))
}
