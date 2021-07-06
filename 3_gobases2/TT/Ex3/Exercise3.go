package main

import (
	"errors"
	"fmt"
)

const (
	SMALL  = "SMALL"
	MEDIUM = "MEDIUM"
	LARGE  = "LARGE"
)

type eCommerce interface {
	price() float64
	shipment() string
}

type store1 struct {
	nombreProd string
	basePrince float64
	address    string
}

func (t1 store1) price() float64 {
	return t1.basePrince
}
func (t1 store1) shipment() string {
	return t1.address
}

type store2 struct {
	nombreProd string
	basePrince float64
	address    string
}

func (t2 store2) price() float64 {
	return t2.basePrince * 1.03
}
func (t2 store2) shipment() string {
	return t2.address
}

type store3 struct {
	nombreProd string
	basePrince float64
	address    string
}

func (t3 store3) price() float64 {
	return (t3.basePrince * 1.06) + 2500.00
}
func (t3 store3) shipment() string {
	return t3.address
}

func main() {

	store1, err := newStore(SMALL)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Store 1 Price", store1.price())
		fmt.Println("Store 1 Shipment", store1.price())
	}

	store2, err := newStore(MEDIUM)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Store 2 Price", store2.price())
		fmt.Println("Store 2 Shipment", store2.price())
	}
}

func newStore(prodSize string) (eCommerce, error) {

	switch prodSize {
	case SMALL:
		return store1{
			nombreProd: "prod 1",
			basePrince: 100.00,
			address:    "calle 1"}, nil
	case MEDIUM:
		return store2{nombreProd: "prod 2",
			basePrince: 200.00,
			address:    "calle 2"}, nil
	case LARGE:
		return store3{nombreProd: "prod 3",
			basePrince: 300.00,
			address:    "calle 3"}, nil
	default:
		return nil, errors.New("ERROR")
	}

}
