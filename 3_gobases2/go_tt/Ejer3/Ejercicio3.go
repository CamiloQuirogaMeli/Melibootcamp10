/*
Varias tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar
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
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	PEQUEÑO = "pequeño"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type Ecommerce interface {
	Precio() float64
	Envio() float64
}

type Product struct {
	name         string
	precio       float64
	tipoProducto string
	costoTotal   float64
}

type TiendaUno struct {
	product  Product
	category string
}

type TiendaDos struct {
	product Product
	peso    string
}

func (t TiendaUno) Precio() float64 {
	t.product = costoProducto(t.product)
	return t.product.costoTotal
}
func (t TiendaUno) Envio() float64 {
	t.product = costoProducto(t.product)
	return t.product.costoTotal
}
func (t TiendaDos) Precio() float64 {
	t.product = costoProducto(t.product)
	return t.product.costoTotal
}
func (t TiendaDos) Envio() float64 {
	t.product = costoProducto(t.product)
	return t.product.costoTotal
}

func nuevaTienda(p Product) Ecommerce {

	switch p.tipoProducto {
	case PEQUEÑO:
		return TiendaUno(p)
	case MEDIANO:

	case GRANDE:

	}
}

func costoProducto(producto Product) Product {
	var calculos, costoEnvio float64
	switch producto.tipoProducto {
	case "Pequeño":
		producto.costoTotal = producto.precio
	case "Mediano":
		calculos = (producto.costoTotal * 3) / 100
		producto.costoTotal = producto.precio + calculos
	case "Grande":
		calculos = (producto.costoTotal * 6) / 100
		costoEnvio = 2500
		producto.costoTotal = producto.precio + calculos + costoEnvio
	}
	return producto
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	producto := Product{}
	fmt.Println("Bienvenido a su tienda virtual, para registrar un producto, por favor ingrese los siguientes datos")
	fmt.Println("Nombre del producto")
	scanner.Scan()
	producto.name = scanner.Text()

	fmt.Println("Precio")
	var price float64
	_, err := fmt.Scanf("%f", &price)

	if err != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	producto.precio = price

	println("¿Que tipo de producto es? (pequeño, mediano o grande)")
	scanner.Scan()
	producto.tipoProducto = scanner.Text()

}
