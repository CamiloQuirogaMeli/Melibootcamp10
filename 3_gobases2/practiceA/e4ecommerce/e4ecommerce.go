package main

import (
	"errors"
	"fmt"
)

const (
	SMALL   = "Chico"
	MEDIUM  = "Mediano"
	BIG     = "Grande"
	FRAGILE = "Fragil"
	SPECIAL = "Especial"
)

type Products struct {
	Name string
	Type string
	Size float64
}

type Flete struct {
	Name     string
	Cost     float64
	Products []Products
}

func main() {

	productOne := Products{Name: "P1", Type: SMALL, Size: 50}
	productDos := Products{Name: "P2", Type: MEDIUM, Size: 500}
	productTres := Products{Name: "P3", Type: BIG, Size: 4000}
	productCuatro := Products{Name: "P4", Type: FRAGILE, Size: 600}
	productCinco := Products{Name: "P5", Type: SPECIAL, Size: 50000}

	fleteUno := Flete{}
	fleteUno.addProduct(productOne, productDos, productTres, productCuatro, productCinco)
	fleteUno.calculteShipping()
	fmt.Println("Flete ", fleteUno)
}

func (f *Flete) addProduct(p ...Products) {

	for _, product := range p {
		f.Products = append(f.Products, product)
	}

}

func (f *Flete) calculteShipping() {
	var subTotal float64
	for _, item := range f.Products {
		aux, err := item.totalSize()
		if err == nil {
			subTotal += aux
		} else {
			// Generar un nuevo shiping con el producto especial
			// Si entra acá puede que tambien sea un producto no categorizado,
			// por lo que hbría que hacer el tratamiento del error para diferenciarlo.
			fmt.Println("Los productos especiales solo se pueden enviar junto a productos especiales...")
			fmt.Println("Generando un nuevo envío para este el producto", item.Name)
		}
	}
	if subTotal > 10000000 {
		fmt.Println("Se sobrepasó el límite de cargas")
	} else {
		f.Cost = subTotal
	}
}

func (p Products) totalSize() (float64, error) {
	switch p.Type {
	case SMALL:
		return p.Size, nil
	case MEDIUM:
		return p.Size * 1.05, nil
	case BIG:
		return p.Size * 1.2, nil
	case FRAGILE:
		return p.Size * 1.75, nil
	case SPECIAL:
		return p.Size, errors.New("Este es un producto especial")
	default:
		return 0, errors.New("No se encuentra el tipo de producto")
	}
}
