package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

type Small struct {
	size float64
}

func (s Small) getTotalSize() float64 {
	return s.size
}

type Medium struct {
	size float64
}

func (m Medium) getTotalSize() float64 {
	return m.size * 1.05
}

type Big struct {
	size float64
}

func (b Big) getTotalSize() float64 {
	return (b.size * 1.2)
}

type Fragile struct {
	size float64
}

func (f Fragile) getTotalSize() float64 {
	return (f.size * 1.75)
}

type Special struct {
	size float64
}

func (s Special) getTotalSize() float64 {
	return s.size
}

type Product interface {
	getTotalSize() float64
}

func main() {
	var option int
	var aux int = 1
	var products []Product
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Agregar Producto al Flete\n2. Calcular envio del flete \n 0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var productType string
			fmt.Println("Ingrese el tipo de producto")
			fmt.Scan(&productType)

			product, err := createProduct(productType)

			if err != nil {
				fmt.Println(err)
				break
			}

			products = append(products, product)

			fmt.Println("El producto ha sido agregado satisfactoriamente")
		case 2:

			response := calculateFreight(products)
			fmt.Println("*****RESPUESTA****\n", response)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func createProduct(productType string) (Product, error) {

	const (
		SMALL   = "chico"
		MEDIUM  = "mediano"
		BIG     = "grande"
		FRAGILE = "fragil"
		SPECIAL = "especial"
	)

	switch productType {
	case SMALL:
		var size float64
		fmt.Println("Ingrese el tamaño del producto")
		fmt.Scan(&size)
		return Small{size}, nil
	case MEDIUM:
		var size float64
		fmt.Println("Ingrese el tamaño del producto")
		fmt.Scan(&size)
		return Medium{size}, nil
	case BIG:
		var size float64
		fmt.Println("Ingrese el tamaño del producto")
		fmt.Scan(&size)
		return Big{size}, nil
	case FRAGILE:
		var size float64
		fmt.Println("Ingrese el tamaño del producto")
		fmt.Scan(&size)
		return Fragile{size}, nil
	case SPECIAL:
		var size float64
		fmt.Println("Ingrese el tamaño del producto")
		fmt.Scan(&size)
		return Special{size}, nil
	default:
		return nil, errors.New("¡¡ERROR!! has ingresado un tipo de producto invalido")
	}
}

func calculateFreight(products []Product) string {
	const CAPACITY float64 = 10000000
	var specialSize float64 = 0.0
	var otherSize float64 = 0.0
	var chain string = ""

	for _, p := range products {
		typeProduct := reflect.TypeOf(p).Name()

		switch typeProduct {
		case "Special":
			specialSize += p.getTotalSize()
		default:
			otherSize += p.getTotalSize()
		}
	}

	if specialSize > 0 {
		amount := math.Ceil(specialSize / CAPACITY)

		chain += "\n Se tiene que realizar " + fmt.Sprintf("%d", int(amount)) + " envios especiales"
	}

	//fmt.Println("specialSize: ", specialSize, "\notherSize: ", otherSize)

	if otherSize > 0 {
		amount := int(otherSize / CAPACITY)
		if amount == 0 {
			amount = 1
		}
		chain += "\n Se tiene que realizar " + fmt.Sprintf("%d", amount) + " envios normales"
	}

	if chain == "" {
		chain = "No hay productos para realizar envios"
	}

	return chain
}
