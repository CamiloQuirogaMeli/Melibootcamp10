package main

import "fmt"

/* Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre de la misma a elección).
Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne una interface “Ecommerce”
Interface Ecommerce:
 - El método “Precio” debe retornar el precio total en base al costo del producto y los adicionales si los hubiera.
 - El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/
const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type Tienda struct {
	Nombre     string
	Producto   Producto
	DirCliente string
}

type Producto struct {
	TipoProducto string
	CostoProduct float64
	CostoEnvio   float64
	CostoManten  float64
}

type Ecommerce interface {
	Precio() float64
	Envio() string
}

func (t Tienda) Precio() float64 {

	pFinal := t.Producto.CostoEnvio + (t.Producto.CostoManten * t.Producto.CostoProduct)
	return pFinal
}

func (t Tienda) Envio() string {
	return t.DirCliente
}

func nuevaTienda(nom string, dir string, tp string, cost float64) Ecommerce {

	switch tp {

	case CHICO:
		return Tienda{Nombre: nom, DirCliente: dir, Producto: Producto{TipoProducto: tp, CostoProduct: cost, CostoEnvio: 0.0, CostoManten: 1.0}}
	case MEDIANO:
		return Tienda{Nombre: nom, DirCliente: dir, Producto: Producto{TipoProducto: tp, CostoProduct: cost, CostoEnvio: 0.0, CostoManten: 1.03}}
	case GRANDE:
		return Tienda{Nombre: nom, DirCliente: dir, Producto: Producto{TipoProducto: tp, CostoProduct: cost, CostoEnvio: 2500.0, CostoManten: 1.06}}

	default:
		return nil
	}

}

func main() {

	t1 := nuevaTienda("t1", "Laleli 234", CHICO, 150.5)
	t2 := nuevaTienda("t2", "Lofesdi 456", GRANDE, 250.5)

	fmt.Printf("Tienda uno abonará: %v \n", t1.Precio())
	fmt.Printf("Tienda dos abonará: %v \n", t2.Precio())
}
