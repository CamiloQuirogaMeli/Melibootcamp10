package main

import (
	"fmt"
)

const (
	smallType  = 1
	mediumType = 2
	bigType    = 3
)

type Store struct {
	Prod    Product
	Address string
}

type Product struct {
	cost  float64
	typ   int16
	price float64
}

type Ecommerce interface {
	Precio() float64
	Envio() string
}

func (t Store) Envio() string {
	return t.Address
}

func (t Store) Precio() float64 {
	switch t.Prod.typ {
	case smallType:
		return float64(t.Prod.cost)
	case mediumType:
		return float64(t.Prod.cost) + (float64(t.Prod.cost) * 0.3)
	case bigType:
		return float64(t.Prod.cost) + (float64(t.Prod.cost) * 0.6) + 2500
	default:
		return 0
	}
}

func newStore(p Product, address string) Ecommerce {
	store := Store{p, address}
	return store
}

func main() {
	med := Product{cost: 100.0, typ: 2}
	big := Product{cost: 200.0, typ: 3}

	var address string
	fmt.Println("Direccion de envío: ")
	fmt.Scanln(&address)
	store1 := newStore(med, address)
	fmt.Printf("El producto mediano vale $%.2f", store1.Precio())
	fmt.Printf("y será entregado en %v \n", store1.Envio())

	fmt.Println("Direccion de envío: ")
	fmt.Scanln(&address)
	store2 := newStore(big, address)
	fmt.Printf("El producto grande vale $%.2f ", store2.Precio())
	fmt.Printf("y será entregado en %v \n", store2.Envio())
}
