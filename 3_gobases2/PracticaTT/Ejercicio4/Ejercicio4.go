// Ejercicio 4 - Envios
// Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
// La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
// Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un adicional al momento de ser enviado:

// Chico: Ningún adicional.
// Mediano: Requiere un %5 más de espacio
// Grande: Requiere un %20 más de espacio
// Frágil: Requiere un %75 más de espacio
// Especial: Sólo puede ser enviado con productos especiales

// Se solicita que:
// Los productos guarden el tamaño y tengan un método Tamaño Total que nos devuelva el espacio en cm3 que requerimos para ser enviado.
// Y una estructura Flete que tenga los métodos:
// Agregar Producto: agregar producto al flete.
// Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede cargar un total de 10.000.000 cm3 por envío.

package main

import (
	"fmt"
	s "strings"
)

const (
	SMALL       = "small"
	BIG         = "big"
	MEDIUM      = "medium"
	FRAGILE     = "fragile"
	SPECIAL     = "special"
	MAXIMUMLOAD = 10000000.00
)

type Product struct {
	typeProduct string
	size        float64
}

type Freight struct {
	products             []Product
	totalSizeShip        float64
	totalSizeShipSpecial float64
}

func (f *Freight) addProduct(p Product) {

	switch p.typeProduct {
	case SPECIAL:
		f.totalSizeShipSpecial += p.tamTotal()
	default:
		f.totalSizeShip += p.tamTotal()
	}
	f.products = append(f.products, p)
}

func (f *Freight) calculateShip() (int, int) {
	var specialShip float64
	var normalShip float64

	specialShip = f.totalSizeShipSpecial / MAXIMUMLOAD
	normalShip = f.totalSizeShip / MAXIMUMLOAD

	if specialShip > 0 && specialShip < 1 {
		specialShip = 1.0
	}
	if normalShip > 0 && normalShip < 1 {
		normalShip = 1.0
	}

	return int(normalShip), int(specialShip)
}

func (p *Product) tamTotal() float64 {
	switch p.typeProduct {
	case SMALL, SPECIAL:
		return p.size
	case BIG:
		return p.size * 1.20
	case MEDIUM:
		return p.size * 1.05
	case FRAGILE:
		return p.size * 1.75
	}
	return 0
}

func main() {

	var enabledProducs = []string{SMALL, BIG, MEDIUM, FRAGILE, SPECIAL}
	var prod Product
	var freight = Freight{
		totalSizeShip:        0,
		totalSizeShipSpecial: 0,
	}
	var cantCm float64 = MAXIMUMLOAD
	var option string = "si"
	var tryAgain bool = true

	fmt.Printf("Los productos especiales solo se pueden llevar con otros productos especiales, por lo que deben haber viajes solo para estods productos\n")
	fmt.Printf("La capacidad total que puede ser ocupada en cm3 por viaje es: %.2f\n", cantCm)
	fmt.Printf("Los productos habilitados son: ")
	for _, p := range enabledProducs {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	for s.Compare(option, "si") == 0 {
		for tryAgain {
			fmt.Printf("Ingrese el tipo de producto a llevar en el flete: ")
			fmt.Scan(&prod.typeProduct)
			for _, ep := range enabledProducs {
				if s.Compare(prod.typeProduct, ep) == 0 {
					tryAgain = false
				}
			}
		}
		tryAgain = true
		for tryAgain {
			fmt.Printf("Ingrese los cm3 que ocupa el producto: ")
			fmt.Scan(&prod.size)
			if prod.size <= 0 {
				fmt.Printf("El valor de los cm3 no puede ser menor/igual a cero")
				tryAgain = true
			} else {
				tryAgain = false
			}
		}
		tryAgain = true
		freight.addProduct(prod)

		fmt.Printf("¿Desea ingresar otro producto?(si,no): ")
		fmt.Scan(&option)
		option = s.ToLower(option)

		for s.Compare(option, "si") != 0 && s.Compare(option, "no") != 0 {
			fmt.Printf("¿Desea ingresar otro producto?(si,no): ")
			fmt.Scan(&option)
			option = s.ToLower(option)
		}
	}

	normalShip, specialShip := freight.calculateShip()

	fmt.Printf("Los viajes que debe hacer para envios normales es: %d\n", normalShip)
	fmt.Printf("Los viajes que debe hacer para envios especiales es: %d\n", specialShip)

}
