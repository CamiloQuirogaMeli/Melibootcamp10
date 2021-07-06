// Ejercicio 3 - Productos
// Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
// Las empresas tienen 3 tipos de productos:
// Pequeño, Mediano y Grande. (Se espera que sean muchos más)
// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

// Sus costos adicionales son:
// Pequeño: El costo del producto (sin costo adicional)
// Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
// Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

// Requerimientos:
// Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre de la misma a elección).
// Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
// Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne una interface “Ecommerce”
// Interface Ecommerce:
//  - El método “Precio” debe retornar el precio total en base al costo del producto y los adicionales si los hubiera.
//  - El método “Envio” debe retornar la dirección de entrega especificada por el cliente.

// No comprendo los requerimientos ya que solo le veo sentido a crear dos estructuras distintas "tiendauno" y "tiendados" si tienen diferentes funcionalidades para el programa

package main

import (
	f "fmt"
)

const (
	BIG        = "big"
	BIGCOST    = 0
	MEDIUM     = "medium"
	MEDIUMCOST = 0.3
	SMALL      = "small"
	SMALLCOST  = 0.6
)

type Ecommerce interface {
	price() float64
	shipment() string
}

type Shop struct {
	name             string
	addresClientShip string
	productSave      string
}

func price() {

}

func newShop(e string, values ...string) Ecommerce {

}

func main() {

	f.Printf("No comprendo los requerimientos ya que no veo sentido armar dos estructuras distintas tiendauno y tienda dos siento que tienen la misma informacióin\n")
	f.Printf("La interface Ecommer le veo sentido si las tiendas tuvieran diferente estructura porque posee diferentes atributos, pero parecen tener lo mismo")

}
