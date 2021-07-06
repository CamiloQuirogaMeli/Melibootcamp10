package main

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	BIG    = "big"
)

type Ecommerce interface {
	price()
	shipping()
}

type Store1 struct {
	Name     string
	Shipping string
	Price    float64
}

type Store2 struct {
	Name     string
	Shipping string
	Price    float64
}

func (s Store1) price() {
	fmt.Println(s.Price)
}

func (s Store1) shipping() {
	fmt.Println(s.Shipping)
}

func (s Store2) price() {
	fmt.Println(s.Price)
}

func (s Store2) shipping() {
	fmt.Println(s.Shipping)
}

func calculateProductPrice(price float64, productType string) float64 {

	switch productType {
	case MEDIUM:
		return price + (price * 0.3)
	case BIG:
		return price + (price * 0.6) + 2500
	}

	return price
}

func newStore(name string, price float64, productType string, shipping string) Ecommerce {

	totalPrice := calculateProductPrice(price, productType)

	if name == "tienda1" {
		return Store1{name, shipping, totalPrice}
	} else {
		return Store2{name, shipping, totalPrice}
	}

	return nil
}

func main() {

	store1 := newStore("tienda1", 17650, SMALL, "casa")
	store2 := newStore("tienda2", 18000, MEDIUM, "trabajo")

	store1.price()
	store1.shipping()

	store2.price()
	store2.shipping()
}
