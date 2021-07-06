// Varias tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar
// productos y retornar el valor del precio total.
// Las empresas tienen 3 tipos de productos:
// - Pequeño, Mediano y Grande. (Se espera que sean muchos más)
// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de
// envío.

// Sus costos adicionales son:
// - Pequeño: El costo del producto (sin costo adicional)
// - Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén
// de la tienda.
// - Grande: El costo del producto + un 6% por mantenimiento, y un costo adicional por
// envío de $2500.

// Requerimientos:
// - Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre
// de la misma a elección).
// - Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
// - Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne
// una interface “Ecommerce”
// - Interface Ecommerce:
// - El método “Precio” debe retornar el precio total en base al costo del producto y los
// adicionales si los hubiera.
// - El método “Envio” debe retornar la dirección de entrega especificada por el cliente.

package main

import "fmt"

//FALTA TERMINAR, PARA MODIFICAR TODO PARA PONER LAS 2 TIENDAS

//declaro mi interfaz y los metodos vacios
type Ecommerce interface {
	precio() float64
	envio() string
}

//genero mis estructuras para implementarles luego los metodos de la interfaz
type tiendaUno struct {
	Nombre    "miTiendaUno"
	Direccion "Paysandu"
}

type tiendaDos struct {
	Nombre    "miTiendaDos"
	Direccion "Salto"
}

//implemento las los metodos dentro de mis tipos de datos que son las tiendas en este caso
func (this tiendaUno) precio() float64 {

}

type Producto struct {
	Precio        float64
	mantenimiento float64
}

type ProductoChico struct {
	Producto
}

func (p ProductoChico) precio() {
	fmt.Println("El precio es $", (p.Precio))
	fmt.Println("Este producto chico no tiene costos adicionales")
}

type ProductoMediano struct {
	Producto
}

func (p ProductoMediano) precio() {
	fmt.Println("El precio total para el producto es $", (p.Precio + p.mantenimiento))
	fmt.Println("Este producto mediano tiene costo de mantenimiento $", p.mantenimiento)
}

type ProductoGrande struct {
	Producto
}

func (p ProductoGrande) precio() {
	fmt.Println("El precio total para el producto es $", (p.Precio + p.mantenimiento))
	fmt.Println("Este producto grande tiene como costo adicional $", p.mantenimiento)
}

const (
	chico   = "C"
	mediano = "M"
	grande  = "G"
)

func main() {
	c := factory(chico, 50000)
	m := factory(mediano, 87000)
	g := factory(grande, 120000)

	c.precio()
	m.precio()
	g.precio()
}

func factory(tipo string, precio float64) producto {
	switch tipo {
	case chico:
		return ProductoChico{Producto{Precio: precio, Adicionales: 0}}
	case mediano:
		return ProductoMediano{Producto{Precio: precio, Adicionales: (precio + (precio * 0.03))}}
	case grande:
		return ProductoGrande{Producto{Precio: precio, Adicionales: (precio + (precio * 0.06) + 2500)}}
	}

	return nil
}
