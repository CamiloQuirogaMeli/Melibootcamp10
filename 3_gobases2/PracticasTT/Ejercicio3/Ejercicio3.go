package main

import "fmt"

const (
	PEQUENO = "pequeno"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

func main() {
	/*
		Ejercicio 3 - Productos
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

	prod1 := nuevaTienda(PEQUENO, 50, "Calle1")
	prod2 := nuevaTienda(MEDIANO, 100, "Calle2")
	prod3 := nuevaTienda(GRANDE, 150, "Calle3")

	fmt.Println("Precio Total:", prod1.precio(), "\nDireccion", prod1.envio())
	fmt.Println("Precio Total:", prod2.precio(), "\nDireccion", prod2.envio())
	fmt.Println("Precio Total:", prod3.precio(), "\nDireccion", prod3.envio())

}

func nuevaTienda(tipoProduto string, precio float64, direccion string) Ecommerce {
	switch tipoProduto {
	case PEQUENO:
		return prodP{precio, direccion}
	case MEDIANO:
		return prodM{precio, direccion}
	case GRANDE:
		return prodG{precio, direccion}
	default:
		return nil
	}
}

type Ecommerce interface {
	precio() float64
	envio() string
}

type prodP struct {
	pr        float64
	direccion string
}

func (p prodP) precio() float64 {
	return p.pr
}

func (p prodP) envio() string {
	return p.direccion
}

type prodM struct {
	pr        float64
	direccion string
}

func (p prodM) precio() float64 {
	return p.pr * 1.03
}

func (p prodM) envio() string {
	return p.direccion
}

type prodG struct {
	pr        float64
	direccion string
}

func (p prodG) precio() float64 {
	return p.pr*1.06 + 2500
}

func (p prodG) envio() string {
	return p.direccion
}
