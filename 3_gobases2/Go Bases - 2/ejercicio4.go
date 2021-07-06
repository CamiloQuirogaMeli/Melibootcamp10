package main

/*
	Ejercicio 4 - Envios
	Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de
	productos:
	La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
	Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto
	requiere un adicional al momento de ser enviado:
		● Chico: Ningún adicional.
		● Mediano: Requiere un %5 más de espacio
		● Grande: Requiere un %20 más de espacio
		● Frágil: Requiere un %75 más de espacio
		● Especial: Sólo puede ser enviado con productos especiales

	Se solicita que:
		1. Los productos guarden el tamaño y tengan un método Tamaño Total que nos
		devuelva el espacio en cm3 que requerimos para ser enviado.
		2. Y una estructura Flete que tenga los métodos:
			- Agregar Producto: agregar producto al flete.
			- Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo
			puede cargar un total de 10.000.000 cm3 por envío.
*/

type Producto struct {
	tipo string
	peso float64
}

type Flete struct {
	tamañoFlete float64
}

func tamañoTotal(tipo string, espacio float64) (string, float64) {
	switch tipo {
	case "Chico":
		return espacio
	case "Mediano":
		return tipo, espacio + (espacio * 0.05)
	case "Grande":
		return tipo, espacio + (espacio * 0.2)
	case "Fragil":
		return tipo, espacio + (espacio * 0.75)
	case "Chico":
		return tipo, espacio
	}
	return "Vacio", 0
}

func (f Flete) AgregarProducto(tipo string, espacio float64, total float64, productosSlice []productos) ([]productos, float64) {
	if total > 10000000 {
		total = total + espacio
		productosSlice = append(productoSlice, productos{tipo: tipo, peso: espacio})
	}
	return productosSlice, total
}

func (f Flete) CalcularEnvios(total float64) {
	cantEnvios := 0
	pesoMax := 10000000
	if total > pesoMax {
		pesoMax = 0
		cantEnvios++
	}
}

func main() {
	var sliceProductos []Producto
	var peso := 0
	producto1 := productosSlice{tipo: "Chico", peso: 5000000}
	producto1 := productosSlice{tipo: "Grande", peso: 500000}
}
