package main

import "fmt"

/*
	Ejercicio 3 - Productos
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
		* El método “Precio” debe retornar el precio total en base al costo del producto y los
		adicionales si los hubiera.
		* El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/

type tiendaUno struct {
	nombre string
	precio float64
}

type tiendaDos struct {
	nombre string
	precio float64
}

type Ecommerce interface {
	Precio(producto string) float64
	Envio(direccion string) string
}

func nuevaTienda(tipo string) Ecommerce {
	if tipo == "tiendaUno" {
		return tiendaUno{nombre: "Samsung", precio: 100000}
	}
	if tipo == "tiendaDos" {
		return tiendaDos{nombre: "LG", precio: 200000}
	}
	return nil
}

func (t1 tiendaUno) Precio(producto string) float64 {
	switch producto {
	case "pequeño":
		return t1.precio
	case "mediano":
		return t1.precio + (t1.precio * 0.03)
	case "grande":
		return t1.precio + (t1.precio * 0.06) + 2500
	}
	return t1.precio
}

func (t1 tiendaUno) Envio(direccion string) string {
	resultado := "se enviara a la direcciòn:" + direccion
	return resultado
}

func (t2 tiendaDos) Precio(producto string) float64 {
	switch producto {
	case "pequeño":
		return t2.precio
	case "mediano":
		return t2.precio + (t2.precio * 0.03)
	case "grande":
		return t2.precio + (t2.precio * 0.06) + 2500
	}
	return t2.precio
}

func (t2 tiendaDos) Envio(direccion string) string {
	resultado := "se enviara a la direcciòn:" + direccion
	return resultado
}

func main() {
	tiendaSamsung := nuevaTienda("tiendaUno")
	precio1 := tiendaSamsung.Precio("mediano")
	envio1 := tiendaSamsung.Envio("Calle 1 # 1 - 1")
	fmt.Printf("La tienda SAMSUNG, entrega el producto con precio %v y %s\n", precio1, envio1)
	tiendaLG := nuevaTienda("tiendaDos")
	precio2 := tiendaLG.Precio("grande")
	envio2 := tiendaLG.Envio("Calle 2 # 2 - 2")
	fmt.Printf("La tienda LG, entrega el producto con precio %v y %s\n", precio2, envio2)
}
