/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos: - Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
- Pequeño: El costo del producto (sin costo adicional)
- Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
- Grande: El costo del producto + un 6% por mantenimiento, y un costo adicional por envío de $2500.
Requerimientos:
- Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre de la misma a elección).
- Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
- Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne una interface “Ecommerce”
- Interface Ecommerce:
- El método “Precio” debe retornar el precio total en base al costo del producto y los adicionales si los hubiera.
- El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/
package main

import "fmt"

const (
	producto_pequeño = "pequeno"
	producto_mediano = "mediano"
	producto_grande  = "grande"
)

type Producto struct {
	tipo           string
	valorProd      float64
	ubicacionEnvio string
}

type Ecommerce interface {
	precio() float64
	envio() string
}

type MELI struct {
	NombreTienda string
	prod         Producto
}

type AMAZON struct {
	NombreTienda string
	prod         Producto
}

func (p Producto) precio() float64 {
	var producto_precio float64
	if p.tipo == producto_pequeño {
		producto_precio = p.valorProd
		fmt.Println("El valor del producto es: ")
	} else if p.tipo == producto_mediano {
		producto_precio = p.valorProd + (p.valorProd * 0.3)
		fmt.Println("El valor del producto es: ")
	} else {
		producto_precio = p.valorProd + (p.valorProd * 0.6) + 2500
		fmt.Println("El valor del producto es: ")
	}
	return producto_precio
}

func (p Producto) envio() string {
	var ubic string = p.ubicacionEnvio
	fmt.Println("La ubicación de envío del producto es: ")
	return ubic
}

func detalle(e Ecommerce) {
	//fmt.Println(e)
	fmt.Println(e.precio())
	fmt.Println(e.envio())
}

func nuevaTienda(tipoProduct string, valor float64, ubic string) Ecommerce {
	switch tipoProduct {
	case producto_pequeño:
		return Producto{tipo: tipoProduct, valorProd: valor, ubicacionEnvio: ubic}
	case producto_mediano:
		return Producto{tipo: tipoProduct, valorProd: valor, ubicacionEnvio: ubic}
	case producto_grande:
		return Producto{tipo: tipoProduct, valorProd: valor, ubicacionEnvio: ubic}
	}
	return nil
}
