package main

import (
	"fmt"
)

const (
	sizeS = "small"
	sizeM = "medium"
	sizeB = "big"
)

var (
	shippingCost = 2500.0
	percentM     = 0.03
	percentB     = 0.06
	localUno     = tiendaUno{"LA UNO", "RIVERA 4563 BIS", 3.5}
	localDos     = tiendaDos{"COMPRE AQUI", "AV. BRASIL 8803", 6.7}
)

type eCommerce interface {
	precio(string) float64
	envio() string
}

type tiendaUno struct {
	name         string
	address      string
	prodBaseCost float64
}

type tiendaDos struct {
	name         string
	address      string
	prodBaseCost float64
}

func main() {

	var shops []eCommerce

	shops = append(shops, newShop("tienda1"))
	shops = append(shops, newShop("tienda3"))
	shops = append(shops, newShop("tienda2"))
	shops = append(shops, newShop("tienda4"))

	for _, value := range shops {
		if value != nil {

			fmt.Println(value.envio())
			fmt.Println(value.precio("small"))
			fmt.Println(value.precio("medium"))
			fmt.Println(value.precio("big"))
		} else {
			fmt.Println("shop not fund")
		}
	}

	fmt.Println()
}

func newShop(shop string) eCommerce {
	switch shop {
	case "tienda1":
		return localUno
	case "tienda2":
		return localDos
	default:
		return nil
	}
}

func (t tiendaUno) precio(prodSize string) float64 {
	return getCost(t.prodBaseCost, prodSize)
}

func (t tiendaUno) envio() string {
	return t.address
}

func (t tiendaDos) precio(prodSize string) float64 {
	return getCost(t.prodBaseCost, prodSize)
}

func (t tiendaDos) envio() string {
	return t.address
}

func getCost(price float64, size string) (cost float64) {

	switch size {
	case sizeS:
		cost = price
	case sizeM:
		cost = price + (price * percentM)
	case sizeB:
		cost = price + (price * percentB) + shippingCost
	default:
		cost = 0.0
	}

	return
}

/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total.

Las empresas tienen 3 tipos de productos:
- Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de
envío.

Sus costos adicionales son:
- Pequeño: El costo del producto (sin costo adicional)
- Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén
de la tienda.
- Grande: El costo del producto + un 6% por mantenimiento, y un costo adicional por
envío de $2500.

Requerimientos:
	- Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre
	de la misma a elección).
	- Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
	- Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne
	una interface “Ecommerce”
	- Interface Ecommerce:
	- El método “Precio” debe retornar el precio total en base al costo del producto y los
	adicionales si los hubiera.
	- El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/
