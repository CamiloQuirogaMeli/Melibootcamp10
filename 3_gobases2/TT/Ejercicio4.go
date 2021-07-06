package main

import (
	"errors"
	"fmt"
)

type product interface {
	totalSize() float64
}

type flete struct {
	products []prod
}

type prod struct {
	size float64
	typeProduct string
}

type smallProduct struct{
	prod
}
type mediumProduct struct{
	prod
}
type largeProduct struct{
	prod
}
type fragileProduct struct{
	prod
}
type specialProduct struct{
	prod
}

const (
	small    = "small"
	medium    = "medium"
	large     = "large"
	SPECIAL = "special"
	FRAGILE = "fragile"
	totalCapacity = 10000000.0
)
func main(){
	flete := flete{}
	flete.addProduct(
		prod{
			size: 1500000.0,
			typeProduct: SPECIAL,
		}, )
	product := largeProduct{
		prod: prod{
			size: 4000000.0,
			typeProduct: large,
		},
	}


	if _, err := flete.addProduct(product.prod); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se agrego correctamente")
	}
	fmt.Println("Envios totales:", flete.totalShippings())

}


func (p prod) totalSize() float64{
	switch p.typeProduct {
	case medium:
		return p.size * 1.05
	case large:
		return p.size * 1.2
	case FRAGILE:
		return p.size * 1.75
	default:
		return p.size
	}
}

func (flet flete) totalShippings() int{

	cap := 0.0
	for _, prod := range flet.products {
		cap += prod.totalSize()
	}
	if cap / totalCapacity < 1.0 {
		return 1
	} else {
		return int(cap / totalCapacity)
	}
}

func (flet *flete) addProduct(p prod) (int, error) {
	if p.size > 10000000.0 {
		return -1, errors.New("producto muy grande")
	}
	for _, pr := range flet.products {
		if (p.typeProduct == SPECIAL && pr.typeProduct != SPECIAL) ||
			(p.typeProduct != SPECIAL && pr.typeProduct == SPECIAL){
			return -1, errors.New("este envio contiene productos no especiales")
		}
	}
	parcialCapacity := p.size

	for _, pr := range flet.products {
		parcialCapacity += pr.totalSize()
		fmt.Println(parcialCapacity > totalCapacity)
		if parcialCapacity > totalCapacity {
			return -1, errors.New("se supero la capacidad maxima del flete")
		}
	}
	flet.products = append(flet.products, p)
	return 0, nil
}
