package main

import (
	"fmt"
)

type Producto interface {
	costoAdicional() float32
}
type Chico struct {
	costo float32
}
type Mediano struct {
	costo float32
}
type Grande struct {
	costo float32
}

func (c Chico) costoAdicional() float32 {
	return c.costo
}
func (c Mediano) costoAdicional() float32 {
	return c.costo * 1.05
}
func (c Grande) costoAdicional() float32 {
	return c.costo * 1.2
}

type Ecommerce interface {
	precio() float32
	envio() string
}
type Tienda struct {
	Prod           Producto
	DireccionEnvio string
}

func (t Tienda) envio() string {
	return t.DireccionEnvio
}
func (t Tienda) precio() float32 {
	return t.Prod.costoAdicional()
}
func newTienda(Prod Producto) Ecommerce {
	tienda := Tienda{Prod, ""}
	fmt.Println("Direccion de envío: ")
	fmt.Scanln(&tienda.DireccionEnvio)
	return tienda
}
func main() {
	mediano := Mediano{costo: 100.0}
	grande := Grande{costo: 200.0}
	tienda1 := newTienda(mediano)
	tienda2 := newTienda(grande)

	fmt.Printf("El producto mediano va a valer $%.2f y va a ser entregado a %s", tienda1.precio(), tienda1.envio())
	fmt.Printf("El producto mediano va a valer $%2.f y va a ser entregado a %s", tienda2.precio(), tienda2.envio())
}
