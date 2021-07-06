package main

import "fmt"

//Un Ecommerce necesita realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene 3 tipos de productos: Chico, Mediano
//y Grande. (Se espera que sean muchos más)

//Y los costos adicionales son:
//- Chico: solo tiene el costo del producto
//- Mediano: el precio del producto + un %3 de mantenerlo en la tienda
//- Grande: el precio del producto + un %6 de mantenerlo en la tienda y adicional a eso $2500 de flete.

//El porcentaje de mantenerlo en la tienda es en base al precio del producto.

//Para ello se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

//Deber poder ejecutar el método Precio y me traiga el precio total en base al costo del producto y los adicionales en caso que los tenga

type producto interface {
	precio()
}

type Producto struct {
	Precio float64
	Adicionales float64
}

type ProductoChico struct {
	Producto
}

func (p ProductoChico) precio(){
	fmt.Println("El precio total para el producto es $", (p.Precio + p.Adicionales))
	fmt.Println("Este producto chico no tiene costos adicionales")
}

type ProductoMediano struct {
	Producto
}

func (p ProductoMediano) precio() {
	fmt.Println("El precio total para el producto es $", (p.Precio + p.Adicionales))
	fmt.Println("Este producto mediano tiene como costo adicional $", p.Adicionales)
}

type ProductoGrande struct {
	Producto
}

func (p ProductoGrande) precio() {
	fmt.Println("El precio total para el producto es $", (p.Precio + p.Adicionales))
	fmt.Println("Este producto grande tiene como costo adicional $", p.Adicionales)
}

const (
	chico = "C"
	mediano = "M"
	grande = "G"
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
		return ProductoChico{Producto{Precio: precio, Adicionales: 0},}
	case mediano:
		return ProductoMediano{Producto{Precio: precio, Adicionales: (precio + (precio * 0.03))},}
	case grande:
		return ProductoGrande{Producto{Precio: precio, Adicionales: (precio + (precio * 0.06) + 2500)},}
	}

	return nil
}