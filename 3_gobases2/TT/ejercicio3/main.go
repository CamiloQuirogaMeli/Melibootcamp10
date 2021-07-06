package main

import (
	"errors"
	"fmt"
)

type Store struct {
	name string
}

func (s Store) getPrice(productType string, productPrice float64) (float64, error) {
	const (
		SMALL  = "chico"
		MEDIUM = "mediano"
		BIG    = "grande"
	)

	switch productType {
	case SMALL:
		return productPrice, nil
	case MEDIUM:
		return productPrice * 1.03, nil
	case BIG:
		return (productPrice * 1.06) + 2500, nil
	default:
		return 0.0, errors.New("¡¡ERROR!! has ingresado un tipo de producto inexistente")
	}
}

func (s Store) getAdress(address string) string {

	return address
}

type Ecommerce interface {
	getPrice(productType string, productPrice float64) (float64, error)
	getAdress(address string) string
}

func main() {
	var option int
	var aux int = 1
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Calcular valor de un producto \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var storeName string
			var address string
			var productType string
			var productPrice float64

			fmt.Println("Ingrese el nombre de la tienda")
			fmt.Scan(&storeName)

			ecommerce := newStore(storeName)
			if ecommerce == nil {
				fmt.Println("Ha ocurrido un error!!")
				break
			}
			fmt.Println("Ingrese su direccion")
			fmt.Scan(&address)

			fmt.Println("Ingrese el tipo de producto")
			fmt.Scan(&productType)

			fmt.Println("Ingrese el valor del producto")
			fmt.Scan(&productPrice)

			totalPrice, err := ecommerce.getPrice(productType, productPrice)

			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("***** RESPUESTA ****\nPrecio Total: ", "$", totalPrice, "\nDireccion: ", ecommerce.getAdress(address))

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func newStore(name string) Ecommerce {

	switch name {
	case "tienda1":
		var t1 Store = Store{name}
		return t1
	case "tienda2":
		var t2 Store = Store{name}
		return t2
	default:
		return nil
	}
}
