package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	flete := Flete{}
	prod, err := newProduct("small", 5000)

	if err == nil {
		flete.AddProduct(prod)
	} else {
		fmt.Println(err)
	}

	prod, err = newProduct("medium", 50000)

	if err == nil {
		flete.AddProduct(prod)
	} else {
		fmt.Println(err)
	}

	prod, err = newProduct("big", 10000000)

	if err == nil {
		flete.AddProduct(prod)
	} else {
		fmt.Println(err)
	}

	prod, err = newProduct("fragile", 2000)

	if err == nil {
		flete.AddProduct(prod)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Cantidad de viajes: ", flete.CalculateShipping())
}

func newProduct(t string, s float64) (Product, error) {

	switch t {
	case "small":
		return SmallProduct{size: s}, nil
	case "medium":
		return MediumProduct{size: s}, nil
	case "big":
		return BigProduct{size: s}, nil
	case "fragile":
		return FragileProduct{size: s}, nil
	case "special":
		return SpecialProduct{size: s}, nil
	default:
		return nil, errors.New("Tipo de producto no registrado")
	}
}

type Product interface {
	TotalSize() float64
}

type SmallProduct struct {
	size float64
}

func (p SmallProduct) TotalSize() float64 {
	return p.size
}

type MediumProduct struct {
	size float64
}

func (p MediumProduct) TotalSize() float64 {
	return p.size + (p.size * 0.05)
}

type BigProduct struct {
	size float64
}

func (p BigProduct) TotalSize() float64 {
	return p.size + (p.size * 0.2)
}

type FragileProduct struct {
	size float64
}

func (p FragileProduct) TotalSize() float64 {
	return p.size + (p.size * 0.75)
}

type SpecialProduct struct {
	size float64
}

func (p SpecialProduct) TotalSize() float64 {
	return 0
}

type Flete struct {
	products []Product
}

func (f *Flete) AddProduct(p Product) {
	f.products = append(f.products, p)
}

func (f Flete) CalculateShipping() int {
	var totalSize float64
	var shippings float64
	for _, i := range f.products {
		totalSize += i.TotalSize()
	}
	shippings = totalSize / 10000000

	return truncate(shippings)
}

func truncate(n float64) int {
	var i = math.Trunc(n)

	if (n - i) > 0 {
		return int(i + 1)
	} else {
		return int(n)
	}
}
