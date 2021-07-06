package main

import "fmt"

const (
	small   = "chico"
	medium  = "mediano"
	large   = "grande"
	fragile = "fragil"
	special = "Especial"
	maxSize = 10000000
)

type product struct {
	size        int
	productType string
}
type Flete struct {
	products      []product
	llevaEspecial bool
}

func (f *Flete) CalcShipping() int {
	totalSizeShipping := 0.0
	for _, singleProduct := range f.products {
		totalSizeShipping += singleProduct.TotalSize()
		if singleProduct.productType == special {
			f.llevaEspecial = true
		}
	}

	return int(totalSizeShipping / maxSize)
}
func (p product) TotalSize() float64 {
	size := float64(p.size)
	switch p.productType {
	case medium:
		return size * 1.05
	case large:
		return size * 1.2
	case fragile:
		return size * 1.75
	default:
		return size
	}

}
func main() {
	fmt.Println("Ejercicio")

}
