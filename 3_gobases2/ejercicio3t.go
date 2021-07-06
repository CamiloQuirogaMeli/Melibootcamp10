package main

// Un Ecommerce necesita realizar una funcionalidad en Go para administrar productos y
// retornar el valor del precio total.
// La empresa tiene 3 tipos de productos: Chico, Mediano y Grande. (Se espera que sean
// muchos más)

// Y los costos adicionales son:
// - Chico: solo tiene el costo del producto
// - Mediano: el precio del producto + un %3 de mantenerlo en la tienda
// - Grande: el precio del producto + un %6 de mantenerlo en la tienda y adicional a eso
// $2500 de flete.

// El porcentaje de mantenerlo en la tienda es en base al precio del producto.

// Para ello se requiere una función factory que reciba el tipo de producto y el precio y retorne
// una interfaz Producto que tenga el método Precio.

// Deber poder ejecutar el método Precio y me traiga el precio total en base al costo del
// producto y los adicionales en caso que los tenga

const (
	chico   = "Chico"
	mediano = "Mediano"
	grande  = "Grande"
)

type Producto interface {
	precio() float64
}

type Chico struct {
	price float64
}

type Mediano struct {
	price float64
}

type Grande struct {
	price float64
}

func (c Chico) precio() float64 {
	return c.price
}
func (m Mediano) precio() float64 {
	return m.price * 1.03
}
func (g Grande) precio() float64 {
	return g.price*1.06 + 2500
}

func factory(op string, price float64) Producto {
	switch op {
	case chico:
		return Chico{price}
	case mediano:
		return Mediano{price}
	case grande:
		return Grande{price}
	default:
		println("No existe este tipo")
		return nil
	}
}
